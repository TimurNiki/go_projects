package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TimurNiki/go_api_tutorial/v4/cmd/api"
	"github.com/TimurNiki/go_api_tutorial/v4/configs"
	"github.com/TimurNiki/go_api_tutorial/v4/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	
	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL!")
}
