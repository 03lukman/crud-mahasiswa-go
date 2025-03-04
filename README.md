# Go CRUD - Aplikasi Manajemen Mahasiswa

## 📌 Deskripsi
Go CRUD adalah aplikasi sederhana berbasis web untuk manajemen data mahasiswa menggunakan **Golang (Go) dengan Gin Framework**, serta MySQL sebagai database. Aplikasi ini memiliki fitur CRUD (Create, Read, Update, Delete) untuk mengelola data mahasiswa.

---

## 🚀 Fitur
- **Menampilkan daftar mahasiswa** dalam tabel responsif.
- **Menambahkan mahasiswa baru** dengan formulir input.
- **Mengedit data mahasiswa** yang sudah ada.
- **Menghapus mahasiswa** dari database.

---

## 🛠️ Teknologi yang Digunakan
- **Golang** (Gin Framework) - Backend API
- **MySQL** - Database
- **Bootstrap** - UI Styling
- **HTML & CSS** - Tampilan frontend

---

## 📂 Struktur Proyek
```
|-- main.go
|-- go.mod
|-- go.sum
|-- controller
|   |-- create_student.go
|   |-- delete_student.go
|   |-- hello_world.go
|   |-- index_student.go
|   |-- update_student.go
|-- database
|   |-- database.go
|-- routes
|   |-- routes.go
|-- views
|   |-- create.html
|   |-- index.html
|   |-- update.html
```

---

## 📌 Instalasi & Menjalankan Aplikasi
### 1️⃣ **Clone Repository**
```bash
git clone https://github.com/03lukman/crud-mahasiswa-go.git
cd crud-mahasiswa-go
```

### 2️⃣ **Setup Database**
Buat database MySQL dengan nama `go_crud_db` dan jalankan SQL berikut:
```sql
CREATE DATABASE go_crud_db;
USE go_crud_db;

CREATE TABLE students (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    nim VARCHAR(20) NOT NULL UNIQUE,
    address TEXT
);
```

### 3️⃣ **Konfigurasi Database**
Edit file `database/database.go` dan sesuaikan dengan koneksi database kamu:
```go
package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	dsn := "root@tcp(localhost:3306)/crud-student"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

```

### 4️⃣ **Jalankan Aplikasi**
```bash
go mod tidy
go run main.go
```
Aplikasi akan berjalan di `http://localhost:8080`.

---

## 🔥 API Endpoints
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/` | Menampilkan daftar mahasiswa |
| GET | `/student/create` | Form tambah mahasiswa |
| POST | `/student/store` | Menyimpan data mahasiswa |
| GET | `/student/update?id=1` | Form edit mahasiswa |
| POST | `/student/update` | Menyimpan perubahan mahasiswa |
| GET | `/student/delete?id=1` | Menghapus mahasiswa |

---

📬 Kontak
Jika Anda memiliki pertanyaan, silakan hubungi saya melalui email: lukmannurhakim3130@gmail.com
---
Terima kasih telah menggunakan Go CRUD! 🎉

