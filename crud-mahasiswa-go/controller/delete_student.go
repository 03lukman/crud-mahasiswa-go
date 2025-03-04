package controller

import (
	"database/sql"
	"net/http"
)

func NewDeleteStudentController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		_, err := db.Exec("DELETE FROM student WHERE id = ?", id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/student", http.StatusMovedPermanently)
	}
}
