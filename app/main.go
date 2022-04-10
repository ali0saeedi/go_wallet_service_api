package main

import (
	_ "github.com/ali0saeedi/go_wallet_service_api/app/database"
	orm "github.com/ali0saeedi/go_wallet_service_api/app/database"
	"github.com/ali0saeedi/go_wallet_service_api/app/router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8000")
}