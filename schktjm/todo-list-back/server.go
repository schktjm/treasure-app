package server

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"firebase.google.com/go/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"github.com/voyagegroup/treasure-app/controller"
	"github.com/voyagegroup/treasure-app/db"
	"github.com/voyagegroup/treasure-app/firebase"
	"github.com/voyagegroup/treasure-app/middleware"
	"github.com/voyagegroup/treasure-app/sample"
)

type Server struct {
	db         *sqlx.DB
	router     *mux.Router
	authClient *auth.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {
	authClient, err := firebase.InitAuthClient()
	if err != nil {
		log.Fatalf("failed init auth client. %s", err)
	}
	s.authClient = authClient

	cs := db.NewDB(datasource)
	dbcon, err := cs.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuth(s.authClient, s.db)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	authChain := commonChain.Append(
		authMiddleware.Handler,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/public").Handler(commonChain.Then(sample.NewPublicHandler()))
	r.Methods(http.MethodGet).Path("/private").Handler(authChain.Then(sample.NewPrivateHandler(s.db)))

	articleController := controller.NewArticle(s.db)
	r.Methods(http.MethodPost).Path("/articles").Handler(authChain.Then(AppHandler{articleController.Create}))
	r.Methods(http.MethodPut).Path("/articles/{id}").Handler(authChain.Then(AppHandler{articleController.Update}))
	r.Methods(http.MethodDelete).Path("/articles/{id}").Handler(authChain.Then(AppHandler{articleController.Destroy}))
	r.Methods(http.MethodGet).Path("/articles").Handler(commonChain.Then(AppHandler{articleController.Index}))
	r.Methods(http.MethodGet).Path("/articles/{id}").Handler(commonChain.Then(AppHandler{articleController.Show}))

	articleCommentController := controller.NewArticleComment(s.db)
	r.Methods(http.MethodPost).Path("/articles/{article_id}/comments").Handler(authChain.Then(AppHandler{articleCommentController.Create}))

	r.PathPrefix("").Handler(commonChain.Then(http.StripPrefix("/img", http.FileServer(http.Dir("./img")))))
	return r
}
