package routes

import (
	"database/sql"
	"net/http"

	"github.com/03lukman/crud-student-go/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/student", controller.NewIndexStudent(db))
	server.HandleFunc("/student/create", controller.NewCreateStudentController(db))
	server.HandleFunc("/student/update", controller.NewUpdateStudentController(db))
	server.HandleFunc("/student/delete", controller.NewDeleteStudentController(db))
}
