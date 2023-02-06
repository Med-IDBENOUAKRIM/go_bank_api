package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	port_address string
	store        Storage
}

func NewServer(port string, store Storage) *ApiServer {
	return &ApiServer{
		port_address: port,
		store:        store,
	}
}

func (s *ApiServer) RunServer() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHttpHandleFunc(s.handleGetAccountById))
	log.Printf("The server it's running on: %+v", s.port_address)
	http.ListenAndServe(s.port_address, router)

}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("Error: %s", r.Method)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return nil
	}

	return toJSON(w, http.StatusOK, accounts)
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	created_account := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(created_account); err != nil {
		return err
	}

	account := NewAccount(created_account.FirstName, created_account.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return toJSON(w, http.StatusCreated, account)
}

func (s *ApiServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid id given %s", idStr)
	}
	account, err := s.store.GetAccountById(id)
	if err != nil {
		return err
	}
	return toJSON(w, http.StatusOK, account)
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func toJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHttpHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			toJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
