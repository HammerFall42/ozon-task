package main

import (
	"fmt"

	"./repository"
)

func main() {

	db, err := repository.NewPostgresCon(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "postgrepass",
	})

	if err == nil {
		udb := repository.NewUrlsDb(db)
		shortened, err := udb.AddNewUrl("some url")
		basic, err := udb.GetUrl(shortened)
		_ = err
		fmt.Print(basic)
	} else {
		fmt.Print(err)
	}

}