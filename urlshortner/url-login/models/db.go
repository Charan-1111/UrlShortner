package models

type DB struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"database"`
}

type Queries struct {
	InsertQueries InsertQueries `json:"insert_queries"`
	SelectQueries SelectQueries `json:"select_queries"`
	UpdateQueries UpdateQueries `json:"update_queries"`
	DeleteQueries DeleteQueries `json:"delete_queries"`
}

type InsertQueries struct {
	InsertUser string `json:"insert_user"`
	InsertJwt  string `json:"insert_jwt"`
}

type SelectQueries struct {
	CheckEmail        string `json:"check_email"`
	CheckUsername     string `json:"check_username"`
	GetHashedPassword string `json:"get_hashed_password"`
	CheckUserJwt      string `json:"check_user_jwt"`
}

type UpdateQueries struct {
}

type DeleteQueries struct {
	DeleteJwt string `json:"delete_jwt"`
}
