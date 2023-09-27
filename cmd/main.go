package main

import (
	"log"

	"github.com/mkock/mytodo/db"

	"github.com/mkock/mytodo/server"
)

func main() {
	db.MustConnect()
	err := server.Serve()
	log.Fatal(err)
}
