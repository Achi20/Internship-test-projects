package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int64          `db:"id"`
	Username string         `db:"username"`
	Password string         `db:"password"`
	Avatar   sql.NullString `db:"avatar"`
}

type Post struct {
	ID          int64     `db:"id"`
	Description string    `db:"description"`
	Image       string    `db:"image"`
	User_ID     int64     `db:"user_id"`
	Created_at  time.Time `db:"created_at"`
}

type Comment struct {
	ID      int64  `db:"id"`
	Post_ID int64  `db:"post_id"`
	User_ID int64  `db:"user_id"`
	Text    string `db:"text"`
}
