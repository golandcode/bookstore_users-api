package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	pq_users_host     = "pq_users_host"
	pq_users_port     = "pq_users_port"
	pq_users_username = "pq_users_username"
	pq_users_password = "pq_users_password"
	pq_users_dbname   = "pq_users_dbname"
)

var (
	Client   *sql.DB
	host     = os.Getenv(pq_users_host)
	port     = os.Getenv(pq_users_port)
	user     = os.Getenv(pq_users_username)
	password = os.Getenv(pq_users_password)
	dbname   = os.Getenv(pq_users_dbname)
)

func init() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	Client, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
