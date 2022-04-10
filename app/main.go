package main

import (
	_ "github.com/ali0saeedi/go_wallet_service_api/database"
	// orm "github.com/ali0saeedi/go_wallet_service_api/database"
	"github.com/ali0saeedi/go_wallet_service_api/router"
)

func main() {
	// defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8000")
}