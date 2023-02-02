package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KwesiLarbi/blog-service/configs"
	"github.com/KwesiLarbi/blog-service/models"
	"github.com/KwesiLarbi/blog-service/responses"
	"github.com/KwesiLarbi/blog-service/helpers"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// VerifyPassword checks the input password while verifying it with the password in the DB
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or password is incorrect")
		check = false
	}

	return check, msg
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
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

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": "Error occured whle checking for email"},
			})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		if count  > 0 {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": "This email already exists!"},
			})
			return
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserId = user.ID.Hex()
		token, refreshToken, _ := GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, user.UserId)
		user.Token = &token
		user.RefreshToken = &refreshToken

		result, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			msg := fmt.Sprintf("User item was not created!")

			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": msg},
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, responses.UserResponse{
			Status: http.StatusCreated,
			Message: "success",
			Data: map[string]interface{}{"data": result},
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": "login or password is incorrect"},
			})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": msg},
			})
			return
		}

		token, refreshToken, _ := GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, foundUser.UserId)

		UpdateAllTokens(token, refreshToken, foundUser.UserId)

		c.JSON(http.StatusOK, foundUser)
	}
}