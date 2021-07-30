package models

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() error {
	stringConn := os.Getenv("DATABASE_URL")
	if len(stringConn) == 0 {
		err := errors.New("the enviroment variable DATABASE_URL_GYN is missing")
		log.Fatal("Error:", err.Error())
		return err
	}
	log.Println("Connecting batabase:", stringConn)

	dbCon, err := sql.Open("postgres", stringConn)
	if err != nil {
		log.Fatalln("Error:", err.Error())
		return err
	}
	dbCon.SetConnMaxLifetime(30000000)
	dbCon.SetMaxIdleConns(1)
	dbCon.SetMaxOpenConns(10)
	DB = dbCon

	log.Println("Connectecd.")
	return nil
}
