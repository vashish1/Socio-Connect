package database

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//User saves users data
type User struct {
	UUID string
	Username     string
	Email        string
	Passwordhash string
}

//Newuser creates a new usertype
func Newuser(a string,b string,c string) User{
u:=User{
	UUID: GenerateUUID(),
	Username: a,
	Email: b,
	Passwordhash: SHA256ofstring(c),
}
return u
}

//Contact stores the contact details of the users.
type Contact struct{
 Name string
 Email string
 Message string
}

//SHA256ofstring is a function which takes a string a reurns its sha256 hashed form
func SHA256ofstring(p string) string {
	h := sha1.New()
	h.Write([]byte(p))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

//GenerateUUID generates a unique id for every user.
func GenerateUUID() string {

	sd := uuid.New()
	return (sd.String())

}

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Collection, *mongo.Client) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	usercollection := client.Database("SC").Collection("User")
	contactcollection := client.Database("SC").Collection("contact")

	return usercollection, contactcollection, client
}

//Insertintouserdb inserts the data into the database
func Insertintouserdb(collection *mongo.Collection, u User) {

	fmt.Println(u.Username)
	insertResult, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}


//Insertintodb inserts the data into the database
func Insertintodb(collection *mongo.Collection, u Contact) {

	fmt.Println(u.Name)
	insertResult, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}