package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/evandrobarbosadosreis/go-jwt/models"
	"github.com/evandrobarbosadosreis/go-jwt/repository"
	"github.com/evandrobarbosadosreis/go-jwt/utils"
	"github.com/golang-jwt/jwt"
)

var secret string

func init() {
	secret = os.Getenv("JWT_SECRET")
}

func (*Controllers) Signup(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var input models.User
		var error models.Error

		json.NewDecoder(r.Body).Decode(&input)

		if input.Email == "" {
			error.Message = "Email is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if input.Password == "" {
			error.Message = "Password is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		repository := repository.UserRepository{}

		user, err := repository.Signup(db, input)

		if err != nil {
			error.Message = "Server error: " + err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.ResponseWithJson(w, user)
	}
}

func (*Controllers) Login(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var input models.User
		var error models.Error

		json.NewDecoder(r.Body).Decode(&input)

		if input.Email == "" {
			error.Message = "Email is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if input.Password == "" {
			error.Message = "Password is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		repository := repository.UserRepository{}

		user, exists, err := repository.Login(db, input)

		if err != nil {
			error.Message = "Server error: " + err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		if !exists {
			error.Message = "Invalid credentials."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		token, err := utils.GenerateToken(user, secret)

		if err != nil {
			log.Panic(err)
		}

		jwt := models.JWT{
			Token: token,
		}

		w.WriteHeader(http.StatusOK)
		utils.ResponseWithJson(w, jwt)
	}
}

func (*Controllers) AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var errorObject models.Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			errorObject.Message = "Invalid token."
			utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}

		authToken := bearerToken[1]

		token, err := jwt.Parse(authToken, func(t *jwt.Token) (interface{}, error) {

			if _, success := t.Method.(*jwt.SigningMethodHMAC); !success { // type assertions

				return nil, fmt.Errorf("there was an error")
			}
			return []byte(secret), nil
		})

		if err != nil {
			errorObject.Message = err.Error()
			utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}

		if !token.Valid {
			errorObject.Message = "Invalid token."
			utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}

		next.ServeHTTP(w, r)
	}
}
