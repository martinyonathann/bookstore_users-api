package users

import (
	"fmt"

	"github.com/martinyonathann/bookstore_users-api/datasource/mysql/users_db"
	"github.com/martinyonathann/bookstore_users-api/logger"
	"github.com/martinyonathann/bookstore_users-api/utils/errors"
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
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); getErr != nil {
		logger.Error("error when trying to get user by ID ", getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

//Save for save users
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare Save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	insertResult, SaveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if SaveErr != nil {
		logger.Error("error when trying to Save User ", SaveErr)
		return errors.NewInternalServerError("database error")
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", SaveErr)
		return errors.NewInternalServerError("database error")
	}
	user.ID = userID
	return nil
}

//Update for update rows in databases
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare Update statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("error when trying to Update data", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

//Delete function is the function for delete users
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare Delete statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("error when trying to Delete by ID", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}
//FindByStatus for find users
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindbyStatus)
	if err != nil {
		logger.Error("error when trying to prepare Find Users statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to Find Users", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("Database error")
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		logger.Error("error when count len of Result", nil)
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return result, nil
}
