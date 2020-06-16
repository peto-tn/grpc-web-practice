package db

import (
	mock_persistence "blog/app/library/test/mock/infrastructure/persistence"
	"context"
	"database/sql"
	"testing"

	"github.com/tj/assert"
)

func newTestDB() *Connection {
	gdb, _, _ := mock_persistence.OpenDBMock()
	return &Connection{
		Writer: gdb,
		Reader: gdb,
	}
}

func TestDB_GetConnection(t *testing.T) {
	t.Parallel()
	t.Run("transaction test", func(t *testing.T) {
		// Mockの準備
		tx, _, err := newTestTransaction()
		assert.NoError(t, err)
		defer tx.db.Close()
		d := newTestDB()
		defer d.Writer.Close()

		// 検証データ
		ctx := context.Background()
		txDB := tx.db.BeginTx(ctx, &sql.TxOptions{})
		ctx = context.WithValue(ctx, &txKey, txDB)

		// 関数実行
		db := d.GetConnection(ctx)
		assert.Exactly(t, txDB, db)
		assert.NotEqual(t, d.Writer, db)
		txDB.Commit()
	})

	t.Run("no transaction test", func(t *testing.T) {
		// Mockの準備
		d := newTestDB()
		defer d.Writer.Close()

		// 検証データ
		ctx := context.Background()

		// 関数実行
		db := d.GetConnection(ctx)
		assert.Exactly(t, d.Writer, db)
	})
}
