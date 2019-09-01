package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllItem(db *sqlx.DB) ([]model.Item, error) {
	a := make([]model.Item, 0)
	if err := db.Select(&a, `SELECT id, body, user_id FROM items`); err != nil {
		return nil, err
	}
	return a, nil
}

func FindItem(db *sqlx.DB, id int64) (*model.Item, error) {
	a := model.Item{}
	if err := db.Get(&a, `
SELECT id, title, body, user_id FROM items WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateItem(db *sqlx.Tx, a *model.Item) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO items ( body, user_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Body, a.UserID)
}

func UpdateItem(db *sqlx.Tx, id int64, a *model.Item) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE items SET  body = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec( a.Body, id)
}

func DestroyItem(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM items WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}
