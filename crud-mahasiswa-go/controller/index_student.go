package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Student struct {
	Id      string
	Name    string
	NIM     string
	Address string
}

func NewIndexStudent(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, nim, address FROM student")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var students []Student
		for rows.Next() {
			var student Student

			err = rows.Scan(
				&student.Id,
				&student.Name,
				&student.NIM,
				&student.Address,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			students = append(students, student)
		}

		fp := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["students"] = students

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}
