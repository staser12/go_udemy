package main

import (
	"fmt"
	"log"
	"net/http"
	"section5/controller"
	"section5/model"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}