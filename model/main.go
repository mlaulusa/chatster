package model

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "chatster"
	password = "chatster"
	database = "chatster"
)

type Model struct {
	Id string `json:"id,omitempty"`
}

var postgres *sql.DB

func init () {

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}

	postgres = db
}

func Close () {
	err := postgres.Close()

	if err != nil {
		log.Fatal(err)
	}
}
