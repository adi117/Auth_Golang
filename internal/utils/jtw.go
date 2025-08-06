package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	Token     string
	ExpiresAt time.Time
	IssuedAt  time.Time
}

func GenerateJWT(userID int, username string) (*TokenResponse, error) {
	expired := time.Now().Add(24 * time.Hour)
	issued := time.Now()

	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expired),
			IssuedAt:  jwt.NewNumericDate(issued),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		Token:     signedToken,
		ExpiresAt: expired,
		IssuedAt:  issued,
	}, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
