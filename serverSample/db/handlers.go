package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	auth "serverSample/auth"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Payload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Mongo
}

type Passcompare struct {
	HashedPass string
	Pass       string
}
type APIError func(http.ResponseWriter, *http.Request) error

func ErrorCathcher(fn APIError) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, "error handler issue", http.StatusInternalServerError)
		}
	})
}

func (p *Payload) PostDataHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	v := new(Payload)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(v)
	HashesPass, err := p.PassHash(v.Password)
	if err != nil {
		log.Println(err)
	}
	v.Password = HashesPass
	if err := p.InsertInDatabase(v); err != nil {
		return fmt.Errorf("error inserting data into database: %+v", err)
	}
	fmt.Println(*v)
	return nil
}

func (p *Payload) LoginHandler(w http.ResponseWriter, r *http.Request) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
		//fmt.Errorf("method not allowed")
	}
	v := new(Payload)
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		log.Fatal(err)
	}
	HashedPass, err := p.PassHash(v.Password)
	if err != nil {
		log.Print("error hashing pass")
	}
	if err := p.Collection.FindOne(context.Background(), bson.M{"email": v.Email}).Decode(v); err != nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(HashedPass), []byte(v.Password)); err != nil {
		http.Error(w, "password incorrect", http.StatusUnauthorized)
		return err
	}
	token, err := auth.GenerateJWT(v.Email, os.Getenv("SECRET"))
	fmt.Println("The Token is:", token)
	if err != nil {
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return err
	}
	if err := json.NewEncoder(w).Encode(token); err != nil {
		log.Fatal("error encoding json token")
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	// Passcompare := Passcompare{
	// 	HashedPass: HashedPass,
	// 	Pass:       v.Password,
	// }
	//log.Println(Passcompare)
	// if err := json.NewEncoder(w).Encode(Passcompare); err != nil {
	// 	log.Fatal("error encoding")
	// }
	return nil
}

func (p *Payload) PassHash(pass string) (string, error) {
	HashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}
	return string(HashedPass), nil
}

func (p *Payload) Handler() {
	http.HandleFunc("/register", ErrorCathcher(p.PostDataHandler))
	http.HandleFunc("/login", ErrorCathcher(p.LoginHandler))
	http.ListenAndServe(":8080", nil)
}
