package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		AllowNativePasswords: true,
		DBName:               "go_api_tutorial",
	}

	sqlStorage := NewMyStorage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)
	api := newAPIServer(":8080", store)
	api.Serve()

}
