package models
import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title   string 						 `json:"title" bson:"title"`
	Content string 						 `json:"content" bson:"content"`
}

