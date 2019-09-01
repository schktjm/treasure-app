package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func CreateArticleComment(db *sqlx.Tx, ac *model.ArticleComment) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article_comment (body, article_id, user_id) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(ac.Body, ac.ArticleID, ac.UserID)
}

func FindCommentsByArticleID(db *sqlx.DB, articleId int64) ([]model.ArticleComment, error) {
	ac := make([]model.ArticleComment, 0)
	if err := db.Select(&ac, `
SELECT id, article_id, user_id, body FROM article_comment WHERE article_id = ?
`, articleId); err != nil {
		return nil, err
	}
	return ac, nil
}
