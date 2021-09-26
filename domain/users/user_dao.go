package users

import (
	"fmt"
	"github.com/babadee08/bookstore_users-api/datasources/mysql/users_db"
	"github.com/babadee08/bookstore_users-api/utils/date_utils"
	"github.com/babadee08/bookstore_users-api/utils/errors"
	"strings"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	indexUniqueEmail = "email_UNIQUE"
	usersDB          = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %v already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %v", err.Error()))
	}

	// Another way is to do this, but it's a little dangerous with worse performance
	// result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %v", err.Error()))
	}
	user.Id = userId

	/*current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registeres", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user*/
	return nil
}
