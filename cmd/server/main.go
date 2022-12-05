package main

import (
	"log"

	"github.com/naruto678/proglog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Println("Starting server at :8080")
	log.Fatal(srv.ListenAndServe())
}
