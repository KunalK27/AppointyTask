package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	U_ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	// Movie   string             `json:"movie,omitempty"`
	// Watched bool               `json:"watched,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
}
type Posts struct {
	P_ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption         string             `json:"caption,omitempty"`
	Imageurl        string             `json:"imageurl,omitempty"`
	PostedTimestamp string             `json:"postedtimestamp,omitempty" `
}
