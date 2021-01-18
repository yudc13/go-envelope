package services

type TransferedStatus int8

const (
	// 转账失败
	TransferedStatusFailure TransferedStatus = -1
	// 余额不足
	TransferedStatusSufficientFunds TransferedStatus = 0
	// 转账成功
	TransferedStatusSuccess TransferedStatus = 1
)

// 转账类型
// 0: 创建账户
// >=1: 进账
// <=-1: 支出
type ChangeType int8

const (
	// 创建账户
	AccountAcreated ChangeType = 0
	// 储值
	AccountStoreValue ChangeType = 1
	// 红包资金支出
	EnvelopeOutgoing ChangeType = -2
	// 红包资金收入
	EnvelopeIncoming ChangeType = 2
	// 红包过期退款
	EnvelopeExpiredRefund ChangeType = 3
)

// 资金交易变化的状态
type ChangeFlag int8

const (
	// 创建账户
	FlagAccountCrated = 0
	// 支出
	FlagTransferOut = -1
	// 收入
	FLagTransferIn = 1
)

// 账户类型
type AccountType int8
const (
	EnvelopeAccount AccountType = 1
)
