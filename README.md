# Portfolio Backend

Backend API untuk website portfolio, dibangun menggunakan Go (Fiber) dan PostgreSQL.

## Tech Stack

- **Language:** Go
- **Framework:** Fiber v2
- **Database:** PostgreSQL
- **Auth:** JWT (dalam pengembangan)

## Struktur Project
```
portfolio-backend/
├── main.go
├── database.go
├── handlers.go
├── education_handlers.go
├── .env
└── go.mod
```

## Setup & Instalasi

1. Clone repository ini
2. Install dependency dengan menjalankan:
```
go mod download
```

3. Bikin file `.env` di root folder, isi dengan:

DB_USER=postgres
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=portfolio_db


4. Jalankan server dengan perintah:
```
go run .
```

5. Server berjalan di `http://localhost:8080`

## API Endpoints

### Projects
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | /projects | Ambil semua project |
| POST | /projects | Tambah project baru |
| PUT | /projects/:id | Update project |
| DELETE | /projects/:id | Hapus project |

### Education
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | /education | Ambil semua data pendidikan |
| POST | /education | Tambah data pendidikan |
| PUT | /education/:id | Update data pendidikan |
| DELETE | /education/:id | Hapus data pendidikan |

## Author

Muhammad Adrian Bintang Hariyanto