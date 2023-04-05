package utils

import (
	"encoding/json"
	"net/http"

	"github.com/evandrobarbosadosreis/go-jwt/models"
	"github.com/golang-jwt/jwt"
)

func RespondWithError(w http.ResponseWriter, statusCode int, error models.Error) {
	w.WriteHeader(statusCode)
	ResponseWithJson(w, error)
}

func ResponseWithJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GenerateToken(user *models.User, tokenSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "jwt-example",
	})
	return token.SignedString([]byte(tokenSecret))
}
