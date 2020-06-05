package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Board struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

var collectionboard *mongo.Collection

func BoardCollection(c *mongo.Database) {
	collectionboard = c.Collection("boards")
}

func GetAllBoards(c *gin.Context) {
	boards := []Board{}
	cursor, err := collectionboard.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all boards, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var board Board
		cursor.Decode(&board)
		boards = append(boards, board)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Boards",
		"data":    boards,
	})
	return
}
