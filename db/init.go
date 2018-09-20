package db

import (
	"os"
	"gopkg.in/mgo.v2"
)

var (
	Sessions *mgo.Collection
)

func init() {
	session, err := mgo.Dial(os.Getenv("MONGO"))
	if err != nil {
		panic(err)
	}

	Sessions = session.DB("blog").C("sessions")

}