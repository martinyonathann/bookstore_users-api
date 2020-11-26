package users

import (
	"github.com/martinyonathann/bookstore_users-api/app/datasource/mysql/users_db"
	"github.com/martinyonathann/bookstore_users-api/utils/date_utils"
	"github.com/martinyonathann/bookstore_users-api/utils/errors"
	"github.com/martinyonathann/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser    = "SELECT * FROM users WHERE id = ?;"
)

//Get for get users
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.ID, getErr.Error()))
	}
	return nil
}

//Save for save users
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, SaveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if SaveErr != nil {
		return mysql_utils.ParseError(SaveErr)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(SaveErr)
	}
	user.ID = userID
	return nil
}
