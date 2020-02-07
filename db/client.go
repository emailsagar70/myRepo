package db

import (
	"database/sql"
	"fmt"
	
	_ "github.com/lib/pq"

	"github.com/tes-svc/config"
)

var (
	db  *sql.DB
	err error
)

// We need to establish a connection with PostGreS database which is hosted in AWS (cloud).
// To establish a connection we need a client. To connect  between AWS(cloud) and laptop we need to write a Client

// NewClient is used to establish a connect between Postgres database and our service.
func NewClient() {
	var err error
	// enable database connections
	db, err = sql.Open(config.DBName, fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.Port, config.DBUserName, config.DBPassword, config.DBName))
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully established DB connection")

}
