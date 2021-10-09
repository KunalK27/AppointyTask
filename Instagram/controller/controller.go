package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kunalk27/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://KunalKulkarni:kunal@cluster0.fs6cl.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "users"
const colName = "userinfo"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb Connection Successful")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

func insertOneUser(insta model.Users) {
	inserted, err := collection.InsertOne(context.Background(), insta)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one User in DB with the ID", inserted.InsertedID)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var uservar model.Users
	_ = json.NewDecoder(r.Body).Decode(&uservar)
	insertOneUser(uservar)
	json.NewEncoder(w).Encode(uservar)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	displayUser(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func displayUser(postId string) {
	id, _ := primitive.ObjectIDFromHex(postId)
	filter := bson.M{"_id": id}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var details []bson.M
	if err = cursor.All(context.Background(), &details); err != nil {
		log.Fatal(err)
	}
	fmt.Println(details)
}
