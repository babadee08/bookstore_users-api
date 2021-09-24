package users

import (
	"github.com/babadee08/bookstore_users-api/domain/users"
	"github.com/babadee08/bookstore_users-api/services"
	"github.com/babadee08/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	/*bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle error
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		// TODO: Handle JSON error
		return
	}*/
	// This function can replace the one above
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
