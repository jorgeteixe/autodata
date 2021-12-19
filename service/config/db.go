package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	var err error
	url := getConnUrl()

	DB, err = sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}
}

func getConnUrl() string {
	godotenv.Load()

	MysqlHost := os.Getenv("MYSQL_HOST")
	MysqlDB := os.Getenv("MYSQL_DATABASE")
	MysqlUser := os.Getenv("MYSQL_USER")
	MysqlPassword := os.Getenv("MYSQL_PASSWORD")
	return MysqlUser + ":" + MysqlPassword + "@tcp(" + MysqlHost + ":3306)/" + MysqlDB
}
