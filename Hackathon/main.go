package main

import (
	"Huy/auth"
	"Huy/file"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	srv := auth.OauthServer()

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) { // not work
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) { // work but not complete
		srv.HandleTokenRequest(w, r)
	})

	http.HandleFunc("/upload", file.Upload) // image upload

	log.Fatal(http.ListenAndServe(":6969", nil))
}
