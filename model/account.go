package model

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Site     string
	Account  string
	Username string
	Avatar   string
	Summary  string
	Counts   int
	Review   int
	Errors   int
}

var AccountModel = Account{}

func (AccountModel *Account) AccountIndex() (err string,account []Account)  {
	db.Limit(50).Find(&account)
	if  len(account) == 0 {
		return "没有数据",nil
	}
	return "",account
}