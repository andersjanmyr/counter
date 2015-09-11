package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Counter interface {
	Inc() error
	Count() (int, error)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port: %s", port)
	r := mux.NewRouter()
	counter := setup()
	r.HandleFunc("/counter", counterHandler(counter))
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":"+port, loggedRouter)
}

func setup() Counter {
	if os.Getenv("REDIS_URL") != "" {
		return NewRedisCounter(os.Getenv("REDIS_URL"))
	} else if os.Getenv("MONGO_URL") != "" {
		return NewMongoCounter(os.Getenv("MONGO_URL"))
	} else if os.Getenv("POSTGRES_URL") != "" {
		return NewPostgresCounter(os.Getenv("POSTGRES_URL"))
	} else {
		return NewMemoryCounter()
	}
}

type MongoCounter struct {
	url string
}

func NewMongoCounter(url string) *MongoCounter {
	return &MongoCounter{url}
}

func (self *MongoCounter) Inc() error {
	return nil
}

func (self *MongoCounter) Count() (int, error) {
	return 0, nil
}

type PostgresCounter struct {
	url string
}

func NewPostgresCounter(url string) *PostgresCounter {
	return &PostgresCounter{url}
}

func (self *PostgresCounter) Inc() error {
	return nil
}

func (self *PostgresCounter) Count() (int, error) {
	return 0, nil
}

type MemoryCounter struct {
	counter int
}

func NewMemoryCounter() *MemoryCounter {
	return &MemoryCounter{}
}

func (self *MemoryCounter) Inc() error {
	self.counter++
	return nil
}

func (self *MemoryCounter) Count() (int, error) {
	return self.counter, nil
}

func counterHandler(counter Counter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			count, err := counter.Count()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(strconv.Itoa(count)))
		} else if r.Method == "POST" {
			err := counter.Inc()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			count, err := counter.Count()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(strconv.Itoa(count)))
		}
	}

}
