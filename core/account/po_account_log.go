package account

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"goEnvelope/services"
	"time"
)

type Log struct {
	Id              int64
	TradeNo         string              // 交易单号
	LogNo           string              // 交易流水编号
	AccountNo       string              // 账户编号
	UserId          string              // 用户编号
	Username        sql.NullString      // 用户名
	TargetAccountNo string              // 交易目标账户编号
	TargetUsername  sql.NullString      // 交易目标用户名
	TargetUserId    string              // 交易目标用户编号
	Amount          decimal.Decimal     // 交易金额
	Balance         decimal.Decimal     // 交易后余额
	ChangeType      services.ChangeType // 交易流水类型
	ChangeFlag      services.ChangeFlag // 交易变化标示
	Status          int                 // 交易状态
	Desc            string              // 交易描述
	CreatedAt       time.Time           // 交易时间
}

func (Log) TableName() string {
	return "account_logs"
}
