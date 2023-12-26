package provider

import "database/sql"

func ConnectDB(url string) (*sql.DB, error) {
	return sql.Open("postgres", url+"?sslmode=disable")
}
