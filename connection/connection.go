package connection

import (
	"errors"
	"log"
	"time"

	"usersapi_go/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*=============================================
=      INFO -    MongoDB Conection            =
=============================================*/

// INFO -- ...
var INFO = &mgo.DialInfo{
	Addrs:    []string{"mongo:27017"},
	Timeout:  60 * time.Second,
	Database: "goDB",
	// ? if you have user and password for access to DB
	Username: "",
	Password: "",
}

// DBNAME - Name For Database Instance
const DBNAME = "goDB"

// DOCNAME - Name of the Doc
const DOCNAME = "Users"
var db *mgo.Database

// COLLECTION - Name for the collection
const COLLECTION = "Users"

/*=============================================
=      Function for Insert a new user         =
=============================================*/

// Insert - Insert a new user inside the DB
func Insert(User model.User) error {
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()

	User.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(User)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

/*=============================================
=      Function for find a new user by ID     =
=============================================*/

// FindByID - Find User by ID inside the DB
func FindByID(id string) (model.User, error) {
	var user model.User
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return user, err
	}

	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	oid := bson.ObjectIdHex(id)
	err = c.FindId(oid).One(&user)
	return user, err
}

/*=============================================
=      Function for Update a user             =
=============================================*/

// Update - Update data user registered in DB yet
func Update(user model.User) error {
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	err = c.UpdateId(user.ID, &user)
	return err
}

/*=============================================
=      Function for find user                 =
=============================================*/

// FindByUser - Find User by ID
func FindByUser(idUser int) ([]model.User, error) {
	var users []model.User
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return users, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	err = c.Find(bson.M{"user_id": idUser}).All(&users)
	return users, err
}

/*=============================================
=      Function for Delete a user             =
=============================================*/

// Delete - Delete user By Obj ID
func Delete(user model.User) error {

	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	err = c.RemoveId(user.ID)

	return err
}