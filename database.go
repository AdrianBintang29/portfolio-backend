package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal membaca file .env:", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Gagal membuka koneksi:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	fmt.Println("Berhasil konek ke database!")
	DB = db
}
