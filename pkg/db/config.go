package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

func NewDBConn() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")
	options := &pg.Options{
		User:     "postgres",
		Password: "",
		Addr:     address,
		Database: "BookStore",
		PoolSize: 50,
	}
	con = pg.Connect(options)
	if con == nil {
		fmt.Println("cannot connect to postgres")
	}
	return
}
