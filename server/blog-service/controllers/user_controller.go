package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/KwesiLarbi/blog-service/configs"
	"github.com/KwesiLarbi/blog-service/models"
	"github.com/KwesiLarbi/blog-service/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status: http.StatusBadRequest, 
				Message: "error", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data": validationErr.Error()},
			})
			return
		}

		newUser := models.User{
			Id: primitive.NewObjectID(),
			Name: user.Name,
			Email: user.Email,
			Password: user.Password,
			CreationTime: time.Now(),
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{
			Status: http.StatusCreated,
			Message: "success",
			Data: map[string]interface{}{"data": result},
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}