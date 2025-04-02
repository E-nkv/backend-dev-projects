package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/E-nkv/backend-dev-projects/restAPI/api"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=learnsql port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return db, err
}

type BasicCredentials struct {
	Nickname *string
	Password *string
}

type User struct {
	ID       int
	Nickname string `gorm:"unique"`
	Password string
}

func basicAuthMdw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var creds BasicCredentials
		if err := dec.Decode(&creds); err != nil || creds.Nickname == nil || creds.Password == nil {
			//occurs when sent credentials are malformed
			api.WriteBadRequestError(w, "invalid credentials")
			return
		}
		if err := IsValidAdmin(*creds.Nickname, *creds.Password); err != nil {
			log.Printf("invalid admin credentials error. User remote_addr: %s\n", r.RemoteAddr)
			api.WriteError(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsValidAdmin(nickname, pass string) error {
	//check in db if nickname exists, and that the hash (most likely bcrypt) of the pass matches with the sent plaintext password.
	var u User
	if res := DB.First(&u, "nickname = ?", nickname); res.Error != nil {
		return res.Error
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	DB = db
	r := chi.NewMux()

	r.With(basicAuthMdw).Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "admin only route!")
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
