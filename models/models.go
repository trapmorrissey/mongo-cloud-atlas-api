package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct

type Utilities struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title,omitempty"`
}
type Config struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn  string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title string             `json:"title" bson:"title,omitempty"`
	Info  *Info              `json:"info" bson:"info,omitempty"`
}
type Info struct {
	Info1 string `json:"info1,omitempty" bson:"info1,omitempty"`
	Info2 string `json:"info2,omitempty" bson:"info2,omitempty"`
	Info3 string `json:"info3,omitempty" bson:"info3,omitempty"`
	Info4 string `json:"info4,omitempty" bson:"info4,omitempty"`
}
