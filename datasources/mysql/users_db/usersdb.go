package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// _ importing MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	mySQLUsersUsername = "mysql_users_username"
	mySQLUsersPassword = "mysql_users_password"
	mySQLUsersHost     = "mysql_users_host"
	mySQLUsersSchema   = "mysql_users_schema"
)

var (
	// ClientDB is the DB connection
	ClientDB *sql.DB

	username = os.Getenv(mySQLUsersUsername)
	password = os.Getenv(mySQLUsersPassword)
	host     = os.Getenv(mySQLUsersHost)
	schema   = os.Getenv(mySQLUsersSchema)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	ClientDB, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = ClientDB.Ping(); err != nil {
		panic(err)
	}

	log.Println("--------------------------------")
	log.Println("Database successfully configured")
	log.Println("--------------------------------")
}
