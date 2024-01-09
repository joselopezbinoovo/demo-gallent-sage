package utils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var PASSWORD = "password"
var SECRET = "thisisasecret"

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println("Cannot load env", e)
	}
	pass := os.Getenv("password")
	secret := os.Getenv("secret")

	if pass != "" {
		PASSWORD = pass
	}

	if secret != "" {
		SECRET = secret
	}
}
func CreateToken(value jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, value)

	tokenString, err := token.SignedString([]byte(SECRET))

	if err != nil {
		panic(err)
	}
	return tokenString
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		fmt.Println("error", err)
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["password"] == PASSWORD
	} else {
		fmt.Println(err)
		return false
	}
}
