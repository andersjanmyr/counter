package main

import (
	"log"
	"strings"

	"gopkg.in/redis.v3"
)

type RedisCounter struct {
	client *redis.Client
	url    string
}

func NewRedisCounter(url string) *RedisCounter {
	log.Printf("Connecting to Redis " + url)
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisCounter{client, url}
}

func (self *RedisCounter) Inc() error {
	log.Printf("Incrementing Redis counter")
	if err := self.client.Incr("counter").Err(); err != nil {
		return err
	}
	return nil
}

func (self *RedisCounter) Count() (int, error) {
	log.Printf("Getting Redis counter")
	cmd := self.client.Get("counter")

	// if counter is nil
	if cmd.Err() != nil && strings.Contains(cmd.Err().Error(), "nil") {
		return 0, nil
	}
	n, err := cmd.Int64()
	if cmd.Err() != nil {
		return 0, err
	}
	return int(n), nil
}
