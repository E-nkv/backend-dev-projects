package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	api "github.com/E-nkv/backend-dev-projects/authentication"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var JWT_SIGNING_SECRET = "super_secret_stuff"
var JWT_EXPIRATION_TIME = time.Hour * 24 * 2

func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=learnsql port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return db, err
}

func withJwtAuthMdw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		if err != nil {
			api.WriteUnauthorizedError(w, "no auth cookie found")
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (any, error) {
			return []byte(JWT_SIGNING_SECRET), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			api.WriteUnauthorizedError(w, "token invalid or expired. please login again "+err.Error())
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp, ok := claims["exp"].(float64)
			if !ok {
				api.WriteUnauthorizedError(w, "missing exp claim in token")
				return
			}
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				api.WriteUnauthorizedError(w, "token expired")
				return
			}
		} else {
			api.WriteUnauthorizedError(w, "invalid token claims")
			return
		}
		next.ServeHTTP(w, r)
	})
}

type LoginPayload struct {
	Nickname string `gorm:"unique"`
	Password string
}
type User struct {
	ID       int
	Nickname string `gorm:"unique"`
	Password string
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		api.WriteBadRequestError(w, "bad request")
		return
	}

	if err := IsValidUser(payload.Nickname, payload.Password); err != nil {
		api.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(JWT_EXPIRATION_TIME).Unix(),
	})
	tokenString, err := token.SignedString([]byte(JWT_SIGNING_SECRET))
	if err != nil {
		api.WriteInternalServerError(w, "error signing the jwt token for authentication. please try again"+err.Error())
		return
	}
	c := http.Cookie{
		Name:     "auth",
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
		Value:    tokenString,
	}
	fmt.Printf("created cookie and sent it to client %+v\n", c)
	http.SetCookie(w, &c)
	api.WriteJSON(w, http.StatusOK, "logged in successfully", "message")
}

func IsValidUser(nick, pass string) error {
	var u User
	if res := DB.First(&u, "nickname = ?", nick); res.Error != nil {
		return res.Error
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
		return err
	}
	return nil
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Value:    "",
		Name:     "auth",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	})
	api.WriteJSON(w, http.StatusOK, "succesfully logged out", "message")
}
func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	DB = db
	mu := chi.NewMux()
	mu.Post("/login", handleLogin)
	mu.With(withJwtAuthMdw).Put("/logout", handleLogout)
	mu.With(withJwtAuthMdw).Get("/protectedRoute", func(w http.ResponseWriter, r *http.Request) {
		api.WriteJSON(w, http.StatusOK, "hi from protected route!", "data")
	})
	log.Fatal(http.ListenAndServe(":8080", mu))
}
