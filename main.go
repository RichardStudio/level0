package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"level0/database"
	"level0/handlers"
	"level0/kafka"
)

func CreateDbConfig() database.DatabaseConfig {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}
	return database.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func main() {
	err := godotenv.Load("config.env")
	client, err := database.NewClient(CreateDbConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer client.DBClose()

	client.RestoreCacheFromDB()

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	go kafka.ConsumeKafkaMessages(client, kafkaBroker, kafkaTopic)

	r := mux.NewRouter()
	r.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
