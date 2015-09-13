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

func NewMongoCounter(url string) (*MongoCounter, error) {
	log.Printf("Connecting to Mongo " + url)
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	c := session.DB("counter-db").C("counter")
	err = c.Insert(bson.M{"n": 0})
	if err != nil {
		return nil, err
	}

	return &MongoCounter{session, url}, nil
}

func (self *MongoCounter) Inc() error {
	c := self.session.DB("counter-db").C("counter")
	err := c.Update(bson.M{}, bson.M{"$inc": bson.M{"n": 1}})
	if err != nil {
		log.Printf("%#v\n", err)
		return err
	}
	return nil
}

func (self *MongoCounter) Count() (int, error) {
	c := self.session.DB("counter-db").C("counter")
	var result interface{}
	err := c.Find(bson.M{}).One(&result)
	if err != nil {
		log.Printf("%#v\n", err)
		return 0, err
	}
	doc := result.(bson.M)
	return doc["n"].(int), nil
}
