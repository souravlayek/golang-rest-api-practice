package utils

import (
	"context"
	"errors"
	"log"

	"github.com/souravlayek/rest-api-tutorial/internal/database"
	"github.com/souravlayek/rest-api-tutorial/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneBook(book model.Book) error {
	var author *model.Author
	if err := database.AuthorCollection.FindOne(context.Background(), bson.M{"firstName": book.Author.FirstName}).Decode(&author); err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	if author == nil {
		book.Author.ID = primitive.NewObjectID()
		_, err := database.AuthorCollection.InsertOne(context.TODO(), book.Author)
		if err != nil {
			return err
		}
	}
	book.Author = author

	_, err := database.BookCollection.InsertOne(context.Background(), book)
	if err != nil {
		return err
	}
	return nil
}

func GetOneBook(isbn string) (*model.Book, error) {
	ctx := context.Background()
	var book model.Book
	cursor := database.BookCollection.FindOne(ctx, bson.M{"isbn": isbn})
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	err := cursor.Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func GetAllBooks() ([]*model.Book, error) {
	ctx := context.TODO()
	cursor, err := database.BookCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var books []*model.Book

	for cursor.Next(ctx) {
		var book *model.Book
		if err = cursor.Decode(&book); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func UpdateBook(isbn string, book model.Book) error {
	ctx := context.Background()
	result, err := database.BookCollection.UpdateOne(ctx, bson.M{"isbn": isbn}, bson.M{"$set": book})
	if err != nil {
		return err
	}
	if result.ModifiedCount != 1 {
		return errors.New("Something Went wrong")
	}
	return nil
}

func DeleteOneBook(isbn string) error {
	ctx := context.Background()
	_, err := database.BookCollection.DeleteOne(ctx, bson.M{"isbn": isbn})
	if err != nil {
		return err
	}
	return nil
}
