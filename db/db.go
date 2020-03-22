package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/toddself/quaranzine/config"
)

var schema = `
CREATE TABLE IF NOT EXISTS author (
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

func buildConnString(cfg *config.Config) string {
	connString := []string{"sslmode=disable"}
	connString = append(connString, fmt.Sprintf("user=%s", cfg.Database.User))
	connString = append(connString, fmt.Sprintf("dbname=%s", cfg.Database.Dbname))
	return strings.Join(connString[:], " ")
}

func Initialize(cfg *config.Config) *sqlx.DB {
	connString := buildConnString(cfg)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
	return db
}
