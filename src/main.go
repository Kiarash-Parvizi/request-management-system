package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	server := Server{router}
	router.HandleFunc("/", server.home).Methods("GET")
	router.HandleFunc("/login", server.login).Methods("GET")
	router.HandleFunc("/login-data", server.loginData).Methods("POST")
	router.HandleFunc("/signup", server.signup).Methods("GET")
	router.HandleFunc("/signup-data", server.signupData).Methods("POST")
	router.HandleFunc("/request-loan", server.requestLoan).Methods("GET")
	router.HandleFunc("/verify-loan", server.verifyLoan).Methods("GET")
	router.HandleFunc("/profile", server.profile).Methods("GET")
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// enforced timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
