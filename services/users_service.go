package services

import (

	"github.com/martinyonathann/bookstore_users-api/domain/users"
	"github.com/martinyonathann/bookstore_users-api/utils/errors"
)

//GetUser function for getUser
func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//CreateUser function for create new user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
