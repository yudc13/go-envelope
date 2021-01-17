package services

import "time"

type AccountService interface {
	// 创建账户
	CreateAccount(dto AccountCreateDTO) (*AccountDTO, error)
	// 转账
	Transfer(dto AccountTransferDTO) (TransferedStatus, error)
	// 储值
	StoreValue(dto AccountTransferDTO) (TransferedStatus, error)
	// 根据用户Id获取红包账户信息
	GetEnvelopeAccountByUserId(userId string) *AccountDTO
}

// 账户交易
type TradeParticipator struct {
	AccountNo string
	UserId    string
	Username  string
}

// 账户转账
type AccountTransferDTO struct {
	TradeNo     string
	TradeBody   TradeParticipator
	TradeTarget TradeParticipator
	AmountStr   string
	ChangeType  int
	ChangeFlag  int
	Desc        string
}

type AccountCreateDTO struct {
	UserId       string
	Username     string
	AccountName  string
	AccountType  ChangeType
	CurrencyType ChangeFlag
	Amount       string
}

type AccountDTO struct {
	AccountCreateDTO
	AccountNo string
	CreateAt  time.Time
}
