package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	address := "127.0.0.1:8080"
	fmt.Println("Server running on port", "http://"+address)

	r := mux.NewRouter()
	r.HandleFunc("/cats", CatsHandler).Methods("GET")

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		Handler:      r,
	}
	log.Fatal(srv.ListenAndServe())
}

func CatsHandler(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("./catdata.json")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	defer file.Close()
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, file)
}
