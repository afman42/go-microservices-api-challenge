package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sesi_3_challenge/routers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "afif123"
	dbname   = "simple-api-go2"
)

func main() {
	PORT := ":8080"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	routers.StartServer(db).Run(PORT)
}
