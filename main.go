package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	render "gopkg.in/unrolled/render.v1"
)

type Counter interface {
	Name() string
	Inc() error
	Count() (int, error)
}

var hostname string
var env string

func main() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		log.Printf("Error getting hostname: %#v", err)
		os.Exit(1)
	}
	env = strings.Join(os.Environ(), "\n")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port: %s", port)
	counter, err := setup()
	if err != nil {
		log.Printf("Error initializing counter: %#v", err)
		os.Exit(1)
	}
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", renderHandler(counter))
	router.HandleFunc("/index.html", renderHandler(counter))
	router.HandleFunc("/counter", counterHandler(counter))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":"+port, loggedRouter)
}

func setup() (Counter, error) {
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

func renderHandler(counter Counter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		render := render.New(render.Options{Layout: "layout"})
		fmt.Printf("%#v", r)
		n, err := counter.Count()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		mountPoint := r.Header.Get("X-Mount-Point")
		if mountPoint == "" {
			mountPoint = os.Getenv("COUNTER_MOUNT_POINT")
		}
		mountPoint = strings.TrimSuffix(mountPoint, "/")

		fmt.Println(mountPoint)
		render.HTML(w, http.StatusOK, "counter",
			map[string]string{
				"count":      strconv.Itoa(n),
				"type":       counter.Name(),
				"Type":       strings.Title(counter.Name()),
				"Hostname":   hostname,
				"Env":        env,
				"MountPoint": mountPoint,
			})
	}
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

type MemoryCounter struct {
	counter int
}

func NewMemoryCounter() (*MemoryCounter, error) {
	return &MemoryCounter{}, nil
}

func (self *MemoryCounter) Name() string {
	return "memory"
}

func (self *MemoryCounter) Inc() error {
	self.counter++
	return nil
}

func (self *MemoryCounter) Count() (int, error) {
	return self.counter, nil
}
