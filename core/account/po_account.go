package account

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

// 账户持久化对象
type Account struct {
	Id           int64           // 账户ID
	AccountNo    string          // 账户编号
	AccountName  string          // 账户名称
	AccountType  int             // 账户类型
	CurrencyCode string          // 货币类型
	UserId       string          // 用户编号
	Username     sql.NullString  // 用户名称
	Balance       decimal.Decimal // 账户可用余额
	Status       int             // 账户状态
	CreatedAt     time.Time       // 创建时间
	UpdatedAt     time.Time       // 更新时间∏
}
