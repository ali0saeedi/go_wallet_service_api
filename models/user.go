package models

import (
	orm "github.com/ali0saeedi/go_wallet_service_api/app/database"
)

type Wallet struct {
	ID      int64 `json:"id"`      // column name `id`
	Balance int64 `json:"balance"` // column name `balance`
	// Password string `json:"password"` // column name `password`
}

var Wallets []Wallet

//Add to
func (wallet Wallet) Insert() (id int64, err error) {

	//adding data
	result := orm.Eloquent.Create(&wallet)
	id = wallet.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// //list
// func (wallet *Wallet) Wallets() (wallets []Wallet, err error) {
// 	if err = orm.Eloquent.Find(&wallets).Error; err != nil {
// 		return
// 	}
// 	return
// }

//get by id
func (wallet *Wallet) GetWalletById(id int64) (requestedWallet Wallet, err error) {

	if err = orm.Eloquent.First(&requestedWallet, id).Error; err != nil {
		// fmt.Println("error", err)
		return
	}
	// fmt.Println("fetched record")
	// fmt.Println(requestedWallet.ID)
	// fmt.Println(requestedWallet.Balance)
	return
}

//Update
func (wallet *Wallet) Update(id int64) (updateWallet Wallet, err error) {

	if err = orm.Eloquent.Select([]string{"id", "balance"}).First(&updateWallet, id).Error; err != nil {
		return
	}

	//parameter1:is the data to be modified
	//parameter2:is the data to be modified
	if err = orm.Eloquent.Model(&updateWallet).Updates(&wallet).Error; err != nil {
		return
	}
	return
}

// //delete data
// func (wallet *Wallet) Destroy(id int64) (Result Wallet, err error) {

// 	if err = orm.Eloquent.Select([]string{"id"}).First(&wallet, id).Error; err != nil {
// 		return
// 	}

// 	if err = orm.Eloquent.Delete(&wallet).Error; err != nil {
// 		return
// 	}
// 	Result = *wallet
// 	return
// }
