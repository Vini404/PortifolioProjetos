package auth

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"secbank.api/dto"
	"strings"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Use a strong secret key stored securely

type Claims struct {
	CustomerID int `json:"customer_id"`
	jwt.StandardClaims
}

func GenerateJWT(customer_id int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Set expiration time
	claims := &Claims{
		CustomerID: customer_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			SetForbidden(w)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			SetForbidden(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetCustomerIDByJwtToken(token string) int {
	tokenJWT, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		return jwtSecret, nil
	})

	claims := tokenJWT.Claims.(*Claims)

	return claims.CustomerID

}

func SetForbidden(res http.ResponseWriter) {
	response := dto.Response{
		Success:    false,
		Timestamp:  time.Now(),
		StatusCode: 403,
	}

	SetResponse(res, response)
}

func SetResponse(res http.ResponseWriter, response dto.Response) {
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(response.StatusCode)

	json.NewEncoder(res).Encode(response)
}
