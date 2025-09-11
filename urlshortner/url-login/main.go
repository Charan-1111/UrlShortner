package main

import (
	"fmt"
	"goapp/database"
	"goapp/routers"
	"goapp/utils"
)

// init is a special function in Go that is called after
// the program has initialized the package's variables.
// It is used here to load the configuration from the
// file specified by the filePath variable and to
// initialize the database connection using the
// database.InitDB() function. Additionally, it is
// used to create the tables in the database, if
// they do not already exist, using the
// database.CreateTables() function.
func init() {
	filePath := "manifest/LOCAL/config.json"
	utils.LoadConfig(filePath)

	database.InitDB()
	database.CreateTables()
}

func main() {
	// start the server
	server := routers.SetUpFiberServer()
	if err := server.Listen(":8000"); err != nil {
		fmt.Println("Failed to start server:", err)
	} else {
		fmt.Println("Server started on port 8080")
	}
}

// TODO: Add graceful shutdown to the server
// TODO: Add logging to the server
// TODO: Add set new password handler
// TODO: Add forgot password handler
