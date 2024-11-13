// jwt_generator.go
package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// Initialize the secret key from environment variables
var jwtSecretKey []byte

//init function runs automatically when the package is imported
func init() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Set jwtSecretKey from environment variable
    jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
    if len(jwtSecretKey) == 0 {
        log.Fatal("JWT_SECRET_KEY is not set in the environment")
    }
}

// GenerateJWT generates and returns a new JWT token.
func GenerateJWT(userID string) (string, error) {
    claims := jwt.MapClaims{
        "userID": userID,
        "exp":    time.Now().Add(time.Hour * 24).Unix(), // 24-hour expiration
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(jwtSecretKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// DecodeJWT decodes the token and returns the userID
func DecodeJWT(tokenString string) (string, error) {
    // Parse and validate the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Ensure the signing method is HMAC (HS256)
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtSecretKey, nil
    })

    if err != nil {
        return "", err
    }

    // Extract and return the userID from token claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if userID, ok := claims["userID"].(string); ok {
            return userID, nil
        }
        return "", errors.New("userID not found in token claims")
    }

    return "", errors.New("invalid token")
}
