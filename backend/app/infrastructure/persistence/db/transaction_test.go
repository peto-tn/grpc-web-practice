package db

import (
	"context"
	"errors"
	"testing"

	testutils "blog/app/library/test"
	mock_persistence "blog/app/library/test/mock/infrastructure/persistence"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newTestTransaction() (*tx, sqlmock.Sqlmock, error) {
	db, mock, err := mock_persistence.OpenDBMock()
	conn := &Connection{
		Writer: db,
		Reader: db,
	}
	return NewTransaction(conn).(*tx), mock, err
}

func TestTransaction_DoInTx(t *testing.T) {
	t.Parallel()
	ctx := testutils.NewTestContext()

	t.Run("commit test", func(t *testing.T) {
		// Mockの準備
		tx, mock, err := newTestTransaction()
		assert.NoError(t, err)
		defer tx.db.Close()

		// 検証データ
		mock.ExpectBegin()
		mock.ExpectCommit()

		// 関数実行
		value, err := tx.DoInTx(ctx, func(context.Context) (interface{}, error) {
			return nil, nil
		})
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.Nil(t, value)
	})

	t.Run("rollback test", func(t *testing.T) {
		// Mockの準備
		tx, mock, err := newTestTransaction()
		assert.NoError(t, err)
		defer tx.db.Close()

		// 検証データ
		expect := "test error"
		mock.ExpectBegin()
		mock.ExpectRollback()

		// 関数実行
		value, err := tx.DoInTx(ctx, func(context.Context) (interface{}, error) {
			return nil, errors.New(expect)
		})
		assert.EqualError(t, err, expect)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.Nil(t, value)
	})

	t.Run("commit error test", func(t *testing.T) {
		// Mockの準備
		tx, mock, err := newTestTransaction()
		assert.NoError(t, err)
		defer tx.db.Close()

		// 検証データ
		expect := "test error"
		mock.ExpectBegin()
		mock.ExpectCommit().WillReturnError(errors.New(expect))
		// mock.ExpectRollback() // FIXME

		// 関数実行
		value, err := tx.DoInTx(ctx, func(context.Context) (interface{}, error) {
			return nil, nil
		})
		assert.EqualError(t, err, expect)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.Nil(t, value)
	})
}
