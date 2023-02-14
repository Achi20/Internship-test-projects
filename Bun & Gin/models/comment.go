package models

type Comment struct {
	ID      int64  `bun:",pk,autoincrement"`
	Post_ID int64  `bun:",notnull"`
	User_ID int64  `bun:",notnull"`
	Text    string `bun:",notnull"`
}
