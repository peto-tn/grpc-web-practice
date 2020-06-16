package db

import (
	"context"
	"database/sql"

	"blog/app/library"

	"github.com/jinzhu/gorm"
)

var txKey = struct{}{}

type tx struct {
	db *gorm.DB
}

func NewTransaction(db *Connection) library.Transaction {
	return &tx{db: db.Writer}
}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (data interface{}, err error) {
	tx := t.db.BeginTx(ctx, &sql.TxOptions{})
	ctx = context.WithValue(ctx, &txKey, tx)

	data, err = f(ctx)
	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
	}
	return
}

func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(&txKey).(*gorm.DB)
	return tx, ok
}
