package main

import (
	"log"
	"v2/cmd/api"
	"v2/db"
	"v2/store"
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

	sqlStorage := db.NewMyStorage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStore(db)
	api := api.NewAPIServer(":8080", store)
	api.Serve()

}
