package main

import (
	"fmt"
	"library/internal/routes"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	port := "8000"
	fmt.Printf("Api running on port %s\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
