package main

import (
	"fmt"
	"net/http"
	"log"
	"./model"
	"./controller"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()

	defer db.Close()

	fmt.Println("Serving...")

	log.Fatal(http.ListenAndServe(":3000", mux))
}