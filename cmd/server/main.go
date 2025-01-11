package main

import (
	"log"

	"github.com/shubham-wawale/dswgo/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
