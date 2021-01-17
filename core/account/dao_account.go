package account

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AccountsDao struct {
	db *gorm.DB
}

// 根据账户编号获取账户信息
func (dao *AccountsDao) GetOne(accountNo string) *Account {
	account := &Account{ AccountNo: accountNo}
	dao.db.Where(account).Find(account)
	return account
}

// 添加一个资金账户
func (dao *AccountsDao) Insert(a *Account) (id int64, err error) {
	tx := dao.db.Create(a)
	return tx.RowsAffected, tx.Error
}
// 根据账户编号和账户类型查询账户信息
func (dao *AccountsDao) GetAccountByUserIdAcoountType(userId string, accountType int) *Account {
	account := &Account{ UserId: userId, AccountType: accountType }
	dao.db.Where(account).Find(account)
	return account
}

// 跟新账户余额
func (dao *AccountsDao) UpdateBalance(accountNo string, balance decimal.Decimal) (id int64, err error) {
	sql := `UPDATE accounts SET balance = balance + CAST(? AS DECIMAL(30, 6)) WHERE account_no = ? and balance >= -1 * CAST(balance AS DECIMAL(30, 6))`
	tx := dao.db.Exec(sql, balance, accountNo, balance)
	return tx.RowsAffected, tx.Error
}


