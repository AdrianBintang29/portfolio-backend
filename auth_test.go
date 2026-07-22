package main

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHashing(t *testing.T) {
	password := "admin12345"

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Gagal hash password: %v", err)
	}

	err = bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		t.Errorf("Password yang benar seharusnya cocok, tapi malah gagal: %v", err)
	}

	err = bcrypt.CompareHashAndPassword(hashed, []byte("password_salah"))
	if err == nil {
		t.Errorf("Password yang salah seharusnya gagal dicocokkan, tapi malah berhasil")
	}
}

func TestJWTGeneration(t *testing.T) {
	claims := jwt.MapClaims{
		"user_id": 1,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		t.Fatalf("Gagal membuat token: %v", err)
	}

	if tokenString == "" {
		t.Errorf("Token yang dihasilkan seharusnya tidak kosong")
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !parsedToken.Valid {
		t.Errorf("Token seharusnya valid ketika diverifikasi ulang, tapi gagal: %v", err)
	}
}

func TestJWTInvalidToken(t *testing.T) {
	invalidToken := "ini.bukan.token.valid"

	_, err := jwt.Parse(invalidToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err == nil {
		t.Errorf("Token yang tidak valid seharusnya menghasilkan error, tapi malah berhasil di-parse")
	}
}
