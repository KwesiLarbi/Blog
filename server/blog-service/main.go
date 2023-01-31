package main

import (
	"github.com/KwesiLarbi/blog-service/configs"
	"github.com/KwesiLarbi/blog-service/routes"
	"github.com/gin-gonic/gin"
)

// tag struct
// type tag struct {
// 	TagID	string	`json:"tag_id"`
// 	Name	string	`json:"name"`
// }

// post struct
// type post struct {
// 	PostID				string 		`json:"post_id"`
// 	UserID				string		`json:"user_id"`
// 	PostTitle			string		`json:"post_title"`
// 	PostText			string		`json:"post_text"`
// 	Tags				[]tag		`json:"tags"`
// 	PostImage			string		`json:"post_image"`
// 	CreationDateTime	time.Time	`json:"creation_date"`
// }

// sample user data
// var users = user{UserID: "3bd82495-98f7-49c1-9785-9b2873fb3a63", Name: "John Hancock", Email: "jhancock@test.com", Password: "123123", AccountCreationDateTime: time.Now()}

// sample post data
// var posts = []post{
// 	{PostID: "01c79f74-759d-477e-a256-a3bd0f32a0ac", UserID: "3bd82495-98f7-49c1-9785-9b2873fb3a63", PostTitle: "First Post!", PostText: "This is my first post! I hope you enjoy my content!", Tags: []tag{{ TagID: "27c0f4cf-c725-44a5-8ba3-89424f156556", Name: "first"}}, PostImage: "some_image_1_url.png", CreationDateTime: time.Now()},
// 	{PostID: "8a590e5e-5f2c-4e74-aa43-c7a8a8ad7125", UserID: "3bd82495-98f7-49c1-9785-9b2873fb3a63", PostTitle: "First Post!", PostText: "This is my first post! I hope you enjoy my content!", Tags: []tag{{ TagID: "5dfab5d4-3e2e-485e-ade4-c2ca2c29a72a", Name: "second"}}, PostImage: "some_image_2_url.png", CreationDateTime: time.Now()},
// 	{PostID: "05c1127b-a045-4cd2-a917-9df89678d4f9", UserID: "3bd82495-98f7-49c1-9785-9b2873fb3a63", PostTitle: "First Post!", PostText: "This is my first post! I hope you enjoy my content!", Tags: []tag{{ TagID: "3e704e1b-7bd1-4625-938c-f6ab0a0fb537", Name: "third"}}, PostImage: "some_image_3_url.png", CreationDateTime: time.Now()},
// }

func main() {
	router := gin.Default()

	// run the database
	configs.ConnectDB()

	// routes
	routes.UserRoute(router)
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	router.Run("localhost:8080")
}