package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"db/db"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am alive!")
}

func initEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func main() {
	initEnv()

	client, err := db.GetMongoClient()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":"+getEnv("PORT"), router))
}