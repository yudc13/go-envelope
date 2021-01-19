package services

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

// 账户服务层接口
type AccountService interface {
	// 创建账户
	CreateAccount(dto AccountCreateDTO) (*AccountDTO, error)
	// 转账
	Transfer(dto AccountTransferDTO) (TransferedStatus, error)
	// 储值
	StoreValue(dto AccountTransferDTO) (TransferedStatus, error)
	GetAccount(accountNo string) (*AccountDTO, error)
	// 根据用户Id获取红包账户信息
	GetEnvelopeAccountByUserId(userId string) (*AccountDTO, error)
}

// 账户交易
type TradeParticipator struct {
	AccountNo string
	UserId    string
	Username  sql.NullString
}

// 账户转账
type AccountTransferDTO struct {
	TradeNo     string
	TradeBody   TradeParticipator
	TradeTarget TradeParticipator
	Amount      decimal.Decimal
	ChangeType  ChangeType
	ChangeFlag  ChangeFlag
	Desc        string
}

type AccountCreateDTO struct {
	UserId       string          `validate:"required"`
	Username     sql.NullString  `validate:"required"`
	AccountName  string          `validate:"required"`
	AccountType  AccountType     `validate:"required"`
	CurrencyCode string          `validate:"required"`
	Amount       decimal.Decimal `validate:"required"` // 账户可用余额
	Status       int             // 账户状态
}

type AccountDTO struct {
	AccountCreateDTO
	AccountNo string
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间∏
}
