package models

var Config Configurations

type Configurations struct {
	Env            string            `json:"env"`
	BcryptCost     int               `json:"bcryptDefaultCost"`
	DBCreds        DB                `json:"db_credentials"`
	DatabaseTables map[string]string `json:"database_tables"`
	Queries        Queries           `json:"queries"`
	Jwt            JWT               `json:"jwt"`
}
