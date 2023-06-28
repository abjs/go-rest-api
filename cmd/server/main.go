package main

import (
	"context"
	"fmt"

	database "github.com/abjs/go-rest-api/internal/db"
)

func Run() error {
	fmt.Println("app starting....")

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the db..")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully connected to database")

	return nil

}

func main() {
	fmt.Println("Go Api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
