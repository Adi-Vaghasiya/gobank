package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("handler error: %v", err)
			if err := WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()}); err != nil {
				log.Printf("Error Writing Json File: %v", err)
			}
		}
	}
}

type APIserver struct {
	listenAddr string
	store      Storage
	Account
}

func NewAPIserver(listenAddr string, store Storage) *APIserver {
	return &APIserver{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIserver) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandleFunc(s.handleGetAccountByID))).Methods("GET")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleDeleteAccount)).Methods("DELETE")
	router.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer)).Methods("GET")
	log.Println("JSON Api Server Started At", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIserver) handleGetAccount(w http.ResponseWriter, _ *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIserver) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		id, error := getID(r)
		if error != nil {
			return error
		}

		getInfoByID, err := s.store.GetAccountByID(id)
		if err != nil {
			return err
		}
		log.Printf("Requested Data For ID: %v\n", id)
		return WriteJSON(w, http.StatusOK, getInfoByID)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method %s not allowed", r.Method)
	// var temp Account
	// err := json.NewDecoder(r.Body).Decode(&temp)
	// if err != nil {
	// 	return err
	// }
	// vars := mux.Vars(r)
	// idStr := vars["id"]
	// id, _ := strconv.Atoi(idStr)
	// start := time.Now()
	// //badData := make(chan int)---------------------(For Test Case IF WriteJson Gives An Error)
	// //return WriteJSON(w, http.StatusOK, badData)--^
	// //account := NewAccount("Admin", "User")
	// //	w.Write([]byte("Test API"))
	// id := mux.Vars(r)["id"]
	// fmt.Println(id)
	// duration := time.Since(start)
	// fmt.Printf("Time Consumed in One Req %v\n", duration.Milliseconds())
	// return WriteJSON(w, http.StatusOK, s.Account)
}

func (s *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		// return err
		return WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Error Encoding Json ",
		})
	}
	account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	tokenString, err := createJWT(account)
	if err != nil {
		return err
	}
	fmt.Println("JWT Token: ", tokenString)
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteAccount(id); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error deleting account: %v", err),
		})
	}
	return WriteJSON(w, http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Account with ID %d deleted successfully", id),
	})
}

func (s *APIserver) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(TransferAccount)
	err := json.NewDecoder(r.Body).Decode(&transferReq)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return WriteJSON(w, http.StatusOK, transferReq)
}

func getID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID: %v", id)
	}
	//, err := s.store.GetAccountByID(id)
	//if err != nil {
	//	return err
	//}
	log.Printf("Requested Data For ID: %v\n", id)
	return id, nil
}

func createJWT(account *Account) (string, error) {
	claims := &jwt.MapClaims{

		"expiresAt":     150000,
		"accountNumber": account.Number,
	}
	secret := os.Getenv("SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validating JWT token")
		tokenString := r.Header.Get("x-jwt-token")
		_, err := validateJWT(tokenString)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Token is Invalid"})
			return
		}
		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}
