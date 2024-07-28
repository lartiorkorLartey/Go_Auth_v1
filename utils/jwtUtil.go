package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ClientClaims struct {
    Email string `json:"email"`
	Role string `json:"role"`
    jwt.RegisteredClaims
}
var jwtKey = []byte(os.Getenv("JWT_KEY"))


func GenerateJWT(email string, role string) (string, error) {
    expirationTime := time.Now().Add(2 * time.Hour) 
    claims := &ClientClaims{
        Email: email,
		Role: role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
	fmt.Println(os.Getenv("JWT_KEY"), err)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ParseJWTWithClaims(tokenString string) (*ClientClaims, error) {
    claims := &ClientClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
        }
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*ClientClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, errors.New("invalid token")
    }
}