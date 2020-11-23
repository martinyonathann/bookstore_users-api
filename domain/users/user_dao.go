package users

import (
	"fmt"

	"github.com/martinyonathann/bookstore_users-api/utils/date_utils"
	"github.com/martinyonathann/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

//Get for get users
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

//Save for save users
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}

	user.DateCreated = date_utils.GetNowString()

	usersDB[user.ID] = user
	return nil
}
