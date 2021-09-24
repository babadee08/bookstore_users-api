package users

import (
	"encoding/json"
	"fmt"
	"github.com/babadee08/bookstore_users-api/domain/users"
	"github.com/babadee08/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle error
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		// TODO: Handle JSON error
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: Handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
