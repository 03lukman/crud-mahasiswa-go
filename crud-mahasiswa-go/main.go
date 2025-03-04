package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/03lukman/crud-student-go/database"
	"github.com/03lukman/crud-student-go/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", server)) // Log error jika server gagal
}
