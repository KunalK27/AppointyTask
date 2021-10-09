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

const connectionStringposts = "mongodb+srv://KunalKulkarni:kunal@cluster0.fs6cl.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbNamepost = "posts"
const colNamepost = "postinfo"

//Most important
var collectionpost *mongo.Collection

//connect with mongodb
func init() {
	//client option
	clientOptions := options.Client().ApplyURI(connectionStringposts)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb Connection Successful")
	collectionpost = client.Database(dbNamepost).Collection(colNamepost)
	//collection instance
	fmt.Println("Collection instance is ready")
}

//MONGODB Helpers -file
//insert 1 record
func insertOnePost(insta model.Posts) {
	inserted, err := collectionpost.InsertOne(context.Background(), insta)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one Post in DB with the ID", inserted.InsertedID)
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var postvar model.Posts
	_ = json.NewDecoder(r.Body).Decode(&postvar)
	insertOnePost(postvar)
	json.NewEncoder(w).Encode(postvar)
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	displayPosts(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func ListPostsofUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	displayPosts(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func displayPosts(postId string) {
	id, _ := primitive.ObjectIDFromHex(postId)
	filter := bson.M{"_id": id}
	cursor, err := collectionpost.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var details []bson.M
	if err = cursor.All(context.Background(), &details); err != nil {
		log.Fatal(err)
	}
	fmt.Println(details)
}
