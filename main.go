package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/leewei05/image-api/rest"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func initDB() {
	_ = godotenv.Load("config.env")

	dbHost := os.Getenv("PG_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbUser := os.Getenv("PG_USER")
	dbPwd := os.Getenv("PG_PWD")
	dbName := os.Getenv("PG_DB")

	pgStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPwd, dbName)

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPwd == "" || dbName == "" {
		log.Panicf("Missing config parameters: %v", pgStr)
	}

	_, err := gorm.Open("postgres", pgStr)
	if err != nil {
		log.Panic("Cannot open PostgreSQL database")
	}
}

func main() {
	initDB()

	serverPort := os.Getenv("HTTP_PORT")
	if serverPort == "" {
		log.Panic("Null HTTP port value")
	}

	r := mux.NewRouter().StrictSlash(true)
	ri := rest.NewRest()

	r.HandleFunc("api/v1/", ri.GetProduct).Methods("GET")
	r.HandleFunc("api/v1/{id}", ri.CreateProduct).Methods("POST")
	r.HandleFunc("api/v1/{id}", ri.UpdateProduct).Methods("PUT")
	r.HandleFunc("api/v1/{id}", ri.DeleteProduct).Methods("DELETE")

	http.Handle("/", r)

	port := fmt.Sprintf(":%v", serverPort)

	s := &http.Server{
		Addr:         port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("HTTP server running on port %v", port)
	log.Fatal(s.ListenAndServe())
}
