package users

import (
	"strings"

	"github.com/martinyonathann/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

//User for model data User
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Auth struct {
	Rc          int64  `json:"rc"`
	Message     string `json:"message"`
	Detail      string `json:"detail"`
	Ext_ref     string `json:"ext_ref"`
	AccessToken string `json:"access_token"`
}

type Users []User

//Validate for validate email
func (user *User) Validate() *errors.RestErr {

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
