package account

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"goEnvelope/services"
	"time"
)

// 账户持久化对象
type Account struct {
	Id           int64                // 账户ID
	AccountNo    string               // 账户编号
	AccountName  string               // 账户名称
	AccountType  services.AccountType // 账户类型
	CurrencyCode string               // 货币类型
	UserId       string               // 用户编号
	Username     sql.NullString       // 用户名称
	Balance      decimal.Decimal      // 账户可用余额
	Status       int                  // 账户状态
	CreatedAt    time.Time            // 创建时间
	UpdatedAt    time.Time            // 更新时间∏
}

func (po *Account) FromDTO(dto *services.AccountDTO) {
	po.AccountNo = dto.AccountNo
	po.AccountName = dto.AccountName
	po.AccountType = dto.AccountType
	po.CurrencyCode = dto.CurrencyCode
	po.UserId = dto.UserId
	po.Username = sql.NullString{
		String: dto.Username,
		Valid: true,
	}
	po.Balance = dto.Amount
	po.Status = dto.Status
	po.CreatedAt = dto.CreatedAt
	po.UpdatedAt = dto.UpdatedAt
}

func (po *Account) ToDTO() *services.AccountDTO {
	dto := &services.AccountDTO{}
	dto.AccountNo = po.AccountNo
	dto.AccountName = po.AccountName
	dto.AccountType = po.AccountType
	dto.CurrencyCode = po.CurrencyCode
	dto.UserId = po.UserId
	dto.Username = po.Username.String
	dto.Amount = po.Balance
	dto.Status = po.Status
	dto.CreatedAt = po.CreatedAt
	dto.UpdatedAt = po.UpdatedAt
	return dto
}
