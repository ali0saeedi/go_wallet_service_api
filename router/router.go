package router

import (
	. "github.com/ali0saeedi/go_wallet_service_api/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// router.GET("/wallets", Wallets)

	// router.GET("/api/v1/wallets/:id", GetWalletById)

	// router.POST("/wallet", Store)

	// router.PUT("/wallet/:id", Update)

	// router.DELETE("/wallet/:id", Destroy)

	//###########
	router.GET("/api/v1/wallets/:id/balance", GetWalletById)
	router.POST("/api/v1/wallets/:id/credit", CreditWalletBalanace)
	router.POST("/api/v1/wallets/:id/debit", DebitWalletBalanace)

	return router
}
