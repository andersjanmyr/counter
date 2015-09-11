package main

import (
	"net/http"
	"strconv"
)
import "os"

type Counter interface {
	Inc() error
	Count() (int, error)
}

func main() {
	counter := setup()
	http.HandleFunc("/counter", counterHandler(counter))
	http.ListenAndServe(":8080", nil)
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
