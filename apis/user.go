package apis

import (
	model "github.com/ali0saeedi/go_wallet_service_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// //list data
// func Wallets(c *gin.Context) {
// 	var wallet model.Wallet
// 	i, err := strconv.ParseInt(c.Request.FormValue("balance"), 10, 64)
//     if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    -1,
// 			"message": "balance should be a number",
// 		})
// 		return
// 	}
// 	wallet.Balance = i
// 	// wallet.Password = c.Request.FormValue("password")
// 	result, err := wallet.Wallets()

// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    -1,
// 			"message": "Sorry no information found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 1,
// 		"data": result,
// 	})
// }

func GetWalletById(c *gin.Context) {
	var wallet model.Wallet
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := wallet.GetWalletById(id)
	if err != nil {
            wallet.ID = id
        	wallet.Insert()
            result = wallet
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

// //adding data
// func Store(c *gin.Context) {
// 	var wallet model.Wallet
//     i, err := strconv.ParseInt(c.Request.FormValue("balance"), 10, 64)
//     if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    -1,
// 			"message": "balance should be a number",
// 		})
// 		return
// 	}
// 	wallet.Balance = i
// 	id, err := wallet.Insert()

// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    -1,
// 			"message": "add failed",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":    1,
// 		"message": "Added successfully",
// 		"data":    id,
// 	})
// }

//credit wallet balance
func CreditWalletBalanace(c *gin.Context) {
	var wallet model.Wallet
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    current_wallet, err := wallet.GetWalletById(id)
	balance, err := strconv.ParseInt(c.Request.FormValue("balance"), 10, 64)
    if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "balance should be a number",
		})
		return
	} else if balance < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "balance should be greater that zero",
		})
		return
	}
    wallet.Balance = current_wallet.Balance + balance
	result, err := wallet.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "fail to edit",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "Successfully modified",
	})
}

//debit wallet balance
func DebitWalletBalanace(c *gin.Context) {
	var wallet model.Wallet
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    current_wallet, err := wallet.GetWalletById(id)
	balance, err := strconv.ParseInt(c.Request.FormValue("balance"), 10, 64)
    if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "balance should be a number",
		})
		return
    } else if balance < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "balance should be greater that zero",
		})
		return
	}
    wallet.Balance = current_wallet.Balance - balance
    if wallet.Balance < 0{
        c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "not enough credit in this wallet",
		})
		return
    }

	result, err := wallet.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "fail to edit",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "Successfully modified",
	})
}

// //delete data
// func Destroy(c *gin.Context) {
// 	var wallet model.Wallet
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	result, err := wallet.Destroy(id)
// 	if err != nil || result.ID == 0 {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    -1,
// 			"message": "failed to delete",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":    1,
// 		"message": "successfully deleted",
// 	})
// }
