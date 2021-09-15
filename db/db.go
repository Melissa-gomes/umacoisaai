package db

import (
	"database/sql"
)

func ConnectWithDataBase() *sql.DB {
	connection := `user=postgres dbname=loja password=Melissa.geovanna17 host=localhost sslmode=disable`
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
