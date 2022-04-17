package setup

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nra_wsl"
	password = "password"
	dbname   = "db_nra_wsl"
	database = "postgres"
)

var ctxConn = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func Connect(database, ctxConn string) (int, error) {

	// Validate parameters but not create conn
	db, err := sql.Open(database, ctxConn)
	if err != nil {
		return 1, errors.New("Connection cannot be opened")
	}
	defer db.Close()

	// Establish the connection
	err = db.Ping()
	if err != nil {
		return 1, errors.New("Connection cannot be created")
	}

	return 0, nil
}
