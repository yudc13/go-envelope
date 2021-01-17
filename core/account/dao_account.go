package account

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AccountsDao struct {
	db *gorm.DB
}

// 根据账户编号获取账户信息
func (dao *AccountsDao) GetOne(accountNo string) (a *Account, err error) {
	var account Account
	err = dao.db.Where("account_no = ?", accountNo).Find(&account).Error
	return &account, err
}

// 添加一个资金账户
func (dao *AccountsDao) Insert(a *Account) (id int64, err error) {
	tx := dao.db.Create(a)
	return tx.RowsAffected, tx.Error
}

// 根据账户编号和账户类型查询账户信息
func (dao *AccountsDao) GetAccountByUserIdAcoountType(userId string, accountType int) (a *Account, err error) {
	var account Account
	err = dao.db.Where("user_id = ? AND account_type = ?", userId, accountType).Find(&account).Error
	return &account, err
}

// 跟新账户余额
func (dao *AccountsDao) UpdateBalance(accountNo string, balance decimal.Decimal) (id int64, err error) {
	sql := `UPDATE accounts SET balance = balance + CAST(? AS DECIMAL(30, 6)) WHERE account_no = ? and balance >= -1 * CAST(balance AS DECIMAL(30, 6))`
	tx := dao.db.Exec(sql, balance, accountNo, balance)
	return tx.RowsAffected, tx.Error
}
