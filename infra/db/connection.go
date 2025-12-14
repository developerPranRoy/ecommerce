package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getConnectionString(cnf config.DBConfig) string {
	conString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.DB_User,
		cnf.DB_Password,
		cnf.DB_Host,
		cnf.DB_Port,
		cnf.DB_Name,
	)
	if !cnf.EnableSSLMode {
		conString += " sslmode=disable"
	}
	return conString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := getConnectionString(*cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	return dbCon, nil
}
