package main

import (
	"fmt"
	"github.com/toddself/quaranzine/config"
	"github.com/toddself/quaranzine/db"
)

func main() {
	cfg := config.Load()
	db := db.Initialize(&cfg)
	fmt.Println("%v", db)
}
