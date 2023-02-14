package models

import (
	"database/sql"
)

type User struct {
	ID       int64          `bun:",pk,autoincrement"`
	Username string         `bun:",notnull"`
	Password string         `bun:",notnull"`
	Avatar   sql.NullString `bun:",nullzero"`
	// Post     []*Post        `bun:"rel:has-many,join:id=user_id"`
	// Comment  *Comment       `bun:"rel:has-many"`
}
