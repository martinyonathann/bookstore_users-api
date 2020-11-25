package users

import (
	"fmt"
	"strings"

	"github.com/martinyonathann/bookstore_users-api/app/datasource/mysql/users_db"
	"github.com/martinyonathann/bookstore_users-api/utils/date_utils"
	"github.com/martinyonathann/bookstore_users-api/utils/errors"
)
const(
	indexUniqueEmail 	= "email_UNIQUE"
	errorNoRows			= "no rows in result set"
	queryInsertUser 	= "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser 		= "SELECT * FROM users WHERE id = ?;"

)
//Get for get users
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	// results, _ := stmt.Query(user.ID)
	// if err != nil{
	// 	return errors.NewInternalServerError(err.Error())
	// }
	// defer results.Close()

	if err := result.Scan(&user.ID,&user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil{
		if strings.Contains(err.Error(), errorNoRows){
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found", user.ID))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", user.ID, err.Error()))
	} 
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	return nil
}

//Save for save users
func (user *User) Save() *errors.RestErr {
	stmt, err  := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail){
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	// result, err := users_db.Client.Exec(queryInsertUser,user.FirstName, user.LastName, user.Email, user.DateCreated)

	userID, err := insertResult.LastInsertId()
	if err !=  nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.ID = userID
	return nil
}
