package surl

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/url"
)

const (
	db         = "surl"
	collection = "urls"
)

var couldNotResolveCollection = errors.New("Could not resolve collection in Mongo DB")

type MgoStore struct {
	session *mgo.Session
}

type document struct {
	key string
	url *url.URL
}

func NewMgoStore(url string) (*MgoStore, error) {
	if s, err := mgo.Dial(url); err != nil {
		return nil, err
	} else {
		return &MgoStore{s}, nil
	}
}

func (self *MgoStore) findCollection() *mgo.Collection {
	return self.session.DB(db).C(collection)
}

func (self *MgoStore) Get(key string) (*url.URL, error) {
	c := self.findCollection()

	if c == nil {
		return nil, couldNotResolveCollection
	}

	result := document{}
	err := c.Find(bson.M{"key": key}).One(&result)
	if err != nil {
		return nil, err
	}

	return result.url, nil
}

func (self *MgoStore) Put(key string, url *url.URL) error {
	c := self.findCollection()

	if c == nil {
		return couldNotResolveCollection
	}

	return c.Insert(&document{key, url})

}
