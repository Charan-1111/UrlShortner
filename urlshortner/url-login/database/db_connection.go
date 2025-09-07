package database

import (
	"fmt"
	"goapp/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
// InitDB is a function to connect to the database and initialize the DB
// connection. It takes the configuration from the models.Config struct and
// uses it to form the DSN string. It then attempts to connect to the database
// using the sqlx library. If the connection is successful, the db connection is
// stored in the DB variable. If the connection fails, an error message is
// printed to the console.
func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		models.Config.DBCreds.Username,
		models.Config.DBCreds.Password,
		models.Config.DBCreds.Host,
		models.Config.DBCreds.Port,
		models.Config.DBCreds.DatabaseName,
	)

	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to database : ", err)
	} else {
		fmt.Println("Successfully connected to database")
	}
}
