package main

import (
	_ "app/database"
	orm "app/database"
	"app/router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8000")
}