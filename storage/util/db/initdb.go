package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbConn *xorm.Engine

func InitDb(name, sourceName string) (err error) {
	DbConn, err = xorm.NewEngine(name, sourceName)
	if err != nil {
		return err
	}
	return nil
}
