package auth

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"secbank.api/dto"
	"strings"
	"time"
)

var jwtSecret = []byte("c6fa045f97ba66935841bf6bfcc5bf4620cd059348ffe1d2ab39991ff50f7b76d0582f903711f94f73af9f4fc30ae12afdb384088ebb5204a04b28b523c7161fd5ee55dba188536cec714c4d043ce0518c51ca0dff9a00cd867a2a6c1ff9166821770f5a92d44e9a6139171d6af631dfda405fd287ec47caabb9432140cace268f76dbe01245ae80711c3eaf9ecfe7fec8f5fa7753fe8b52e0033340146c3e9374c78c5c8e6630c982681fe034bed8ab83fb65299f30373598d28821ebcf11ff094e1301bea72b4b380d8b804c2cfec3f07527210e61147faf9ab2df5736bdc0635791c69e5619da1328626ad44986133385f9fb6132ab8d25b7c863f250b239") // Use a strong secret key stored securely

// Claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Function to generate JWT token
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Set expiration time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Create the token using HS256 algorithm and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Middleware to verify JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			SetForbidden(w)
			return
		}

		// Extract token from the header (Bearer token)
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		claims := &Claims{}
		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// Check for errors and token validity
		if err != nil || !token.Valid {
			SetForbidden(w)
			return
		}

		// Set the claims as context or session variables if needed
		next.ServeHTTP(w, r)
	})
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
