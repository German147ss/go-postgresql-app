package main

import (
	"os"

	"log"

	"github.com/orlmonteverde/go-postgres-microblog/internal/data"
	"github.com/orlmonteverde/go-postgres-microblog/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// connection to the database.
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Server is listening on port %s", port)

	// start the server.
	serv.Start()
}
