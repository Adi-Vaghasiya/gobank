package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(Str, secret string) (string, error) {
	cwd, _ := os.Getwd()
	fmt.Println("currrent dir is:", cwd)
	// err := godotenv.Load(filepath.Join(cwd, ".env"))
	// if err != nil {
	// 	return "", err
	// }
	claims := jwt.MapClaims{
		"user_email": Str,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("secret is:", secret)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal("error converting the token to string")
	}
	return signedToken, nil
}
