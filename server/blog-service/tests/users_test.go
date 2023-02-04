package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KwesiLarbi/blog-service/models"

	"github.com/stretchr/testify/assert"
)

// test user
var (
	firstName 	= "Ronald"
	lastName	= "Darby"
	password 	= "darby2023"
	email 		= "rdarby@gmail.com"
)


func TestRegister(t *testing.T) {
	newUser := models.User{
		FirstName: 	&firstName,
		LastName: 	&lastName,
		Password: 	&password,
		Email: 		&email,
	}
}

func TestLogin(t *testing.T) {
	fmt.Println("Hello login test")
}