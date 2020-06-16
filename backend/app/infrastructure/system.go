package infrastructure

import (
	"context"
	"errors"

	_ "github.com/jinzhu/gorm/dialects/mysql" // driver

	"blog/app/domain/repository"
	"blog/app/infrastructure/persistence/db"
)

type systemRepository struct {
	conn *db.Connection
}

func NewSystemRepository(conn *db.Connection) repository.SystemRepository {
	return &systemRepository{conn}
}

func (r *systemRepository) Ready(ctx context.Context) error {
	// DB接続チェック
	db := r.conn.Writer
	if db == nil || len(db.GetErrors()) > 0 {
		return errors.New("writer db connection error.")
	}
	db = r.conn.Reader
	if db == nil || len(db.GetErrors()) > 0 {
		return errors.New("reader db connection error.")
	}

	return nil
}
