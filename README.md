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
|-- config
|   |-- database.go
|-- controllers
|   |-- studentController.go
|-- models
|   |-- student.go
|-- routes
|   |-- routes.go
|-- views
|   |-- index.html
|   |-- create.html
|   |-- update.html
|-- public
|   |-- css
|   |-- js
|-- go.mod
|-- go.sum
|-- README.md
```

---

## 📌 Instalasi & Menjalankan Aplikasi
### 1️⃣ **Clone Repository**
```bash
git clone https://github.com/username/go-crud.git
cd go-crud
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
Edit file `config/database.go` dan sesuaikan dengan koneksi database kamu:
```go
package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "root:@tcp(127.0.0.1:3306)/go_crud_db?charset=utf8mb4&parseTime=True&loc=Local"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Gagal koneksi ke database")
    }
    database.AutoMigrate(&models.Student{})
    DB = database
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
Jika Anda memiliki pertanyaan, silakan hubungi saya di:
Email: lukmannurhakim3130@gmail.com
Terima kasih telah menggunakan Go CRUD! 🎉

