package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	ListenAddr string
	store      Storage
	User
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func WriteJson(w http.ResponseWriter, Status int, v any) error {
	w.Header().Set("Content-Type", "appliication/json")
	w.WriteHeader(Status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

func MakeHTTPHandlefunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("handler error: %v", err)
			if err := WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()}); err != nil {
				log.Printf("Error Writing Json File: %v", err)
			}
		}

	}
}

func NewApi(listenaddr string, store Storage) *API {
	return &API{
		ListenAddr: listenaddr,
		store:      store,
	}
}

func (s *API) Run() {
	router := mux.NewRouter()
	//http.HandleFunc("/login", MakeHTTPHandlefunc(s.getAccount))
	router.HandleFunc("/login", MakeHTTPHandlefunc(s.PostUser)).Methods("POST")
	http.ListenAndServe(s.ListenAddr, router)
}

func (s *API) getAccount(w http.ResponseWriter, r *http.Request) error {
	v := "You hit the Login endpoint"
	return WriteJson(w, http.StatusOK, v)
}

func (s *API) PostUser(w http.ResponseWriter, r *http.Request) error {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return err
	}
	user = NewUser(user.FirstName, user.LastName)
	s.store.CreateAccount(user)
	return WriteJson(w, http.StatusOK, user)
}

func PostUser(w http.ResponseWriter, r http.Request) error {
	return nil
}
