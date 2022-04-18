package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

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

func Hello(s string) string {
	return "Hello " + s
}

func GetData() ([]string, error) {

	resultSet := make([]string, 0)
	var s string

	db, err := sql.Open(database, ctxConn)
	if err != nil {
		return make([]string, 1), errors.New("Connection not established")
	}
	defer db.Close()

	rows, err := db.Query("SELECT projectkey, repositoryname FROM result_set")
	if err != nil {
		return make([]string, 1), errors.New("Error in selection")
	}
	defer rows.Close()
	for rows.Next() {
		var projectKey string
		var repoName string
		err = rows.Scan(&projectKey, &repoName)
		if err != nil {
			return make([]string, 1), errors.New("Error in individual row")
		}

		s = fmt.Sprintf("PK: %s, Repo: %s", projectKey, repoName)
		resultSet = append(resultSet, s)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return make([]string, 1), errors.New("Error happened during iteration")
	}

	return resultSet, nil

}

func GetInfoProject(projectName string) error {

	url := fmt.Sprintf("https://api.github.com/users/%s/repos", projectName)
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	//body, err := ioutil.ReadAll(res.Body)
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		msg := fmt.Sprintf("Not successful status: %d", res.StatusCode)
		return errors.New(msg)
	}

	//fmt.Printf("Body: %s\n", body)

	return nil
}
