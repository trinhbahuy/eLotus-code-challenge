package main

import (
	"learn/file"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/upload", file.Upload)

	log.Fatal(http.ListenAndServe(":6969", nil))
}
