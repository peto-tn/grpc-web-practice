package db

import (
	"context"
	"fmt"
	"time"

	"blog/app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Connection struct {
	Writer *gorm.DB
	Reader *gorm.DB
}

func NewDB(conf *config.Config) (*Connection, error) {
	// Writerの接続設定
	wProtocol := fmt.Sprintf("tcp(%s:%s)", conf.DB.Writer.Host, conf.DB.Writer.Port)
	wConnect := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", conf.DB.Writer.User, conf.DB.Writer.Pass, wProtocol, conf.DB.Writer.DBName)
	w, err := gorm.Open(conf.DB.Writer.DBMS, wConnect)
	if err != nil {
		return nil, err
	}
	w.DB().SetConnMaxLifetime(time.Minute * 5)
	w.DB().SetMaxOpenConns(conf.DB.Writer.MaxConnections)

	// Readerの接続設定
	rProtocol := fmt.Sprintf("tcp(%s:%s)", conf.DB.Reader.Host, conf.DB.Reader.Port)
	rConnect := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", conf.DB.Reader.User, conf.DB.Reader.Pass, rProtocol, conf.DB.Reader.DBName)
	r, err := gorm.Open(conf.DB.Reader.DBMS, rConnect)
	if err != nil {
		return nil, err
	}
	r.DB().SetConnMaxLifetime(time.Minute * 5)
	r.DB().SetMaxOpenConns(conf.DB.Reader.MaxConnections)

	return &Connection{
		Writer: w,
		Reader: r,
	}, nil
}

func (c *Connection) GetConnection(ctx context.Context) *gorm.DB {
	conn, ok := GetTx(ctx)
	if !ok {
		return c.Reader
	}
	return conn
}
