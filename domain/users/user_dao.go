package users

import (
	"fmt"

	usersdb "github.com/sasa-radovanovic/bookstore_users-api/datasources/mysql/users_db"
	dateutils "github.com/sasa-radovanovic/bookstore_users-api/utils/date_utils"
	"github.com/sasa-radovanovic/bookstore_users-api/utils/errors"
	mysqlutils "github.com/sasa-radovanovic/bookstore_users-api/utils/mysql_utils"
)

const (
	uniqueEmail       = "email_UNIQUE"
	noRowsInResultSet = "no rows in result set"
	insertQuery       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	selectUserQuery   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	updateUserQuery   = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	deleteUserQuery   = "DELETE FROM users WHERE id=?;"
	findUserByStatus  = "SELECT id, first_name, last_name, email, date_created, status FROM users where status=?;"
)

// Get user from database
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.ClientDB.Prepare(selectUserQuery)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return mysqlutils.ParseMySQLError(err)
	}
	return nil
}

// Save saves user
func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.ClientDB.Prepare(insertQuery)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	user.DateCreated = dateutils.GetNowDBFormat()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		return mysqlutils.ParseMySQLError(err)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlutils.ParseMySQLError(err)
	}
	user.ID = userID
	return nil
}

// Update updates a user
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.ClientDB.Prepare(updateUserQuery)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysqlutils.ParseMySQLError(err)
	}
	return nil
}

// Delete deletes user from the database
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.ClientDB.Prepare(deleteUserQuery)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return mysqlutils.ParseMySQLError(err)
	}
	return nil

}

// FindByStatus finds all users by status
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.ClientDB.Prepare(findUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysqlutils.ParseMySQLError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
