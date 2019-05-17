package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "kamratposten")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "kamratposten")

	password := r.FormValue("password")

	hash, err := hashPassword(password)
	check(err)

	if checkPasswordHash(users["test_user"], hash) {
		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Save(r, w)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "kamratposten")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
