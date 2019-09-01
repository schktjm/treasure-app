package model

type RequestCreateItem struct {
	Body   string  `json:"body"`
}

type ResponseCreateItem struct {
	ID     int64   `json:"id"`
	Body   string  `json:"body"`
	UserID *int64  `json:"user_id"`
	TagIDs []int64 `json:"tag_ids"`
}

type Item struct {
	ID     int64  `db:"id" json:"id"`
	Body   string `db:"body" json:"body"`
	UserID *int64 `db:"user_id" json:"user_id"`
}

type ItemDetail struct {
	Item
}
