package auth

import (
	"fmt"
	"log"
	"testing"
)

func TestAuth(t *testing.T) {
	token, err := GenerateJWT("sample_test", "JWT_TOKEN1976")
	if err != nil {
		log.Fatal("could not generate token")
	}
	fmt.Println(token)
}
