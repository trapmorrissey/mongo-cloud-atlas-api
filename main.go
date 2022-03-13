package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoApi/helper"
	"mongoApi/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Build a FE for PUT/DELETE/GET
// TODO: make live on ubuntu

//Defining logger format
type logWriter struct {
}

//Defining logger details
func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("["+"15:04:05.99Z"+"] - ") + "[api] - " + string(bytes))
}

//Connection mongoDB with helper class
var configCollection = helper.ConnectConfigDB()

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	log.Println("Successfully connected to mongodb")
	//Router
	r := mux.NewRouter()

	// arrange route
	r.HandleFunc("/api/YOUR_ROUTE", createConfig).Methods("POST")
	r.HandleFunc("/api/YOUR_ROUTE", getConfig).Methods("GET")
	r.HandleFunc("/api/YOUR_ROUTE/{id}", updateConfig).Methods("PUT")
	r.HandleFunc("/api/YOUR_ROUTE/{id}", deleteConfig).Methods("DELETE")
	r.HandleFunc("/api/YOUR_ROUTE{id}", specficConfig).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	log.Println("Finished")
	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

func createConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var config models.Config

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&config)

	// insert our book model.
	result, err := configCollection.InsertOne(context.TODO(), config)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
	log.Println("Successfully added new config: " + config.Title)
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var configs []models.Config

	cur, err := configCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var config models.Config
		err := cur.Decode(&config) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		configs = append(configs, config)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(configs) // encode similar to serialize process.
	log.Println("Successfully got current configs")
}

func updateConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var config models.Config

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&config)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"isbn", config.Isbn},
			{"title", config.Title},
			{"info", bson.D{
				{"info1", config.Info.Info1},
				{"info2", config.Info.Info2},
				{"info3", config.Info.Info3},
				{"info3", config.Info.Info4},
			}},
		}},
	}

	err := configCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&config)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	config.ID = id

	json.NewEncoder(w).Encode(config)
	log.Println("Successfully updated: " + config.Title)
}

func deleteConfig(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := configCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func specficConfig(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var config models.Config
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := configCollection.FindOne(context.TODO(), filter).Decode(&config)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(config)
}
