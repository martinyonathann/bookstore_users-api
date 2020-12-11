package users

import (
	"fmt"

	"github.com/martinyonathann/bookstore_users-api/app/datasource/mysql/users_db"
	"github.com/martinyonathann/bookstore_users-api/utils/errors"
	"github.com/martinyonathann/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser      = "SELECT * FROM users WHERE id = ?;"
	queryUpdateUser   = "UPDATE users SET first_name = ? , last_name = ? , email = ? WHERE id = ? ;"
	queryDeleteUser   = "DELETE FROM users WHERE id = ? ;"
	queryFindbyStatus = "SELECT * FROM users WHERE status = ? ;"
)

//Get for get users
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); getErr != nil {
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
	insertResult, SaveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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

//Update for update rows in databases
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindbyStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())

	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())

	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return result, nil
}
