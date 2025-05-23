package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNote(note Note) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return NotesCollection.InsertOne(ctx, note)
}

func GetAllNotes() ([]Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := NotesCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notes []Note
	for cursor.Next(ctx) {
		var note Note
		if err := cursor.Decode(&note); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func UpdateNote(id string, updatedNote Note) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":   updatedNote.Title,
			"content": updatedNote.Content,
		},
	}
	_, err = NotesCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func DeleteNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = NotesCollection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func CreateUser(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := UsersCollection.InsertOne(ctx, user)
	return err
}

func FindUserByUsername(username string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := UsersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}
