package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE author (
	name text,
	email text,
	token text,
	state text
)
`

type Author struct {
	Name  string `db:"name", json:"name"`
	Email string `db:"email", json:"email"`
	Token string `db:"token"`
	State string `db:"state"`
}

func main() {
	db, err := sqlx.Connect("postgres", "user=quaranzine dbname=quaranzine sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

}
