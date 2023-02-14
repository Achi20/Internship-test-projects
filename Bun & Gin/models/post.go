package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Post struct {
	ID          int64  `bun:",pk,autoincrement"`
	Description string `bun:",notnull"`
	Image       string `bun:",notnull"`
	User_ID     int64  `bun:",notnull"`
	// User        *User     `bun:"rel:belongs-to,join:user_id=id"`
	Created_at time.Time `bun:",notnull"`
	// Comment     *Comment  `bun:"rel:has-many"`
}

var _ bun.BeforeAppendModelHook = (*Post)(nil)

func (u *Post) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.Created_at = time.Now()
	}
	return nil
}
