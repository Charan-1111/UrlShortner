package database

import (
	"fmt"
	"goapp/models"
)

func CreateTables() {
	// Implementation to create necessary tables in the database
	for tableName, createQuery := range models.Config.DatabaseTables {
		_, err := DB.Exec(createQuery)
		if err != nil {
			fmt.Printf("Error creating table %s: %v\n", tableName, err)
		} else {
			fmt.Printf("Table %s created or already exists.\n", tableName)
		}
	}
}
