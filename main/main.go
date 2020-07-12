package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/leewei05/image-api/rest"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	db  *gorm.DB
	rdb *redis.Client
)

func initPostgres() {
	_ = godotenv.Load("../config.env")

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

func initRedis() {
	_ = godotenv.Load("../config.env")

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisStr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisStr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Panic("Cannot open Redis")
	}
}

func main() {
	initPostgres()
	initRedis()

	serverPort := os.Getenv("HTTP_PORT")
	if serverPort == "" {
		log.Panic("Null HTTP port value")
	}

	r := mux.NewRouter().StrictSlash(true)
	ri := rest.NewRest(db, rdb)

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
