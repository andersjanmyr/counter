package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoCounter struct {
	session *mgo.Session
	url     string
}

func (self *MongoCounter) Name() string {
	return "mongo"
}

func NewMongoCounter(url string) *MongoCounter {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	c := session.DB("counter-db").C("counter")
	err = c.Insert(bson.M{"n": 0})
	if err != nil {
		panic(err)
	}

	return &MongoCounter{session, url}
}

func (self *MongoCounter) Inc() error {
	c := self.session.DB("counter-db").C("counter")
	err := c.Update(bson.M{"n": 1}, bson.M{"$inc": bson.M{"n": 1}})
	if err != nil {
		log.Printf("%#v\n", err)
		return err
	}
	return nil
}

func (self *MongoCounter) Count() (int, error) {
	c := self.session.DB("counter-db").C("counter")
	var result interface{}
	err := c.Find(bson.M{"n": 1}).One(&result)
	if err != nil {
		log.Printf("%#v\n", err)
		return 0, err
	}
	n := result.(int)
	return n, nil
}
