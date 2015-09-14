package main

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type PostgresCounter struct {
	db  *sql.DB
	url string
}

func NewPostgresCounter(url string) (*PostgresCounter, error) {
	url = url + "?sslmode=disable"
	log.Printf("Connecting to Postgres %s\n", url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Printf("%#v", err)
		return nil, err
	}
	_, err = db.Exec("create table Counter(n integer default 0)")
	if err == nil {
		_, err := db.Exec("insert into Counter(n) values (0)")
		if err != nil {
			log.Printf("%#v", err)
			return nil, err
		}
	} else {
		log.Printf("%#v", err)
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code != "42P07" {
				log.Printf("%#v", err)
				return nil, err
			}
		} else {
			log.Printf("%#v", err)
			return nil, err
		}
	}

	return &PostgresCounter{db, url}, nil
}

func (self *PostgresCounter) Name() string {
	return "postgres"
}

func (self *PostgresCounter) Inc() error {
	_, err := self.db.Exec("update Counter set n = n+1")
	if err != nil {
		log.Printf("%#v", err)
		return err
	}
	return nil
}

func (self *PostgresCounter) Count() (int, error) {
	var n int
	err := self.db.QueryRow("select n from Counter").Scan(&n)
	if err != nil {
		log.Printf("%#v", err)
		return 0, err
	}
	return n, nil
}
