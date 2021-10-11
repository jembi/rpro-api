package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "197.249.6.113"
	port     = 5432
	user     = "postgres"
	password = "temba"
	dbname   = "temba"
)

func GetConnection(ssl bool) *sql.DB {

	var sslMode string

	if ssl == true {
		sslMode = "enabled"
	} else {
		sslMode = "disable"
	}

	connectinString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslMode)

	connection, err := sql.Open("postgres", connectinString)

	if err != nil {
		panic(err)
	}

	err = connection.Ping()
	if err != nil {
		panic(err)
	}

	return connection

}
