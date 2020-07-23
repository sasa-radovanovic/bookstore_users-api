package mysqlutils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sasa-radovanovic/bookstore_users-api/utils/errors"
)

const (
	// NoRows contstant
	NoRows = "no rows in result set"
)

// ParseMySQLError constructs a rest error from database error
func ParseMySQLError(err error) *errors.RestErr {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		fmt.Println(err)
		if strings.Contains(err.Error(), NoRows) {
			return errors.NewNotFoundError("no record matching id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlError.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("invalid data"))
	}
	return errors.NewInternalServerError("error processing request")
}
