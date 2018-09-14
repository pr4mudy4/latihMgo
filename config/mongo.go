package config

import (
	"os"
	mgo "gopkg.in/mgo.v2"

)
func GetMongoDB()(*mgo.Database, error){
	host := os.Getenv("MONGO_HOST")
	user := os.Getenv("MONGO_USER")
	password :=os.Getenv("MONGO_PASSWORD")
	dbName :=os.Getenv("MONGO_DB_NAME")

	session, err := mgo.Dial(host)
	session.DB(dbName).Login(user,password)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}