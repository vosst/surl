package surl

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/url"
)

const (
	collection = "urls"
	db         = "surl"
)

// Indicates that there was an issue resolving the required collection
// from the connected MongoDB instance.
var ErrCouldNotResolveCollection = errors.New("Could not resolve collection")

// MgoStore relies on MongoDB to persist key to URL mappings.
type MgoStore struct {
	session *mgo.Session
}

type Document struct {
	Key string
	Url *url.URL
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
		return nil, ErrCouldNotResolveCollection
	}

	result := Document{}

	err := c.Find(bson.M{"key": key}).One(&result)
	if err != nil {
		return nil, err
	}

	return result.Url, nil
}

func (self *MgoStore) Put(key string, url *url.URL) error {
	c := self.findCollection()

	if c == nil {
		return ErrCouldNotResolveCollection
	}

	return c.Insert(&Document{Key: key, Url: url})

}
