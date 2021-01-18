package account

import (
	"errors"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"goEnvelope/infra/base"
	"goEnvelope/services"
	"gorm.io/gorm"
)

// 真正实现业务逻辑的地方

type domain struct {
	account    Account
	accountLog Log
}

// 创建流水编号logNo
func (d *domain) createAccountLogNo() {
	d.accountLog.LogNo = ksuid.New().Next().String()
}

// 创建账户编号 accountNo
func (d *domain) createAccountNo() {
	d.account.AccountNo = ksuid.New().Next().String()
}

// 创建流水记录
func (d *domain) createAccountLog() {
	// 需要先创建账户
	d.accountLog = Log{}
	// 生成账户流水id
	d.createAccountLogNo()
	d.accountLog.TradeNo = d.accountLog.LogNo
	// 流水中的交易主体信息
	d.accountLog.AccountNo = d.account.AccountNo
	d.accountLog.Username = d.account.Username
	d.accountLog.UserId = d.account.UserId
	// 交易对象
	d.accountLog.TargetAccountNo = d.account.AccountNo
	d.accountLog.TargetUsername = d.account.Username
	d.accountLog.TargetUserId = d.account.UserId
	// 交易金额
	d.accountLog.Amount = d.account.Balance
	d.accountLog.Balance = d.account.Balance
	// 交易变化的属性
	d.accountLog.Desc = "账户创建"
	d.accountLog.ChangeType = services.AccountAcreated
	d.accountLog.ChangeFlag = services.FlagAccountCrated
}
// 创建账户
func (d *domain) Create(dto services.AccountDTO) (*services.AccountDTO, error) {
	// 账户持久化对象
	d.account = Account{}
	// 根据dto初始化account
	d.account.FromDTO(&dto)
	// 生成账户编号
	d.createAccountNo()
	d.account.Username.Valid = true
	// 账户流水持久化对象
	d.createAccountLog()
	accountDao := AccountsDao{}
	accountLogDao := LogDao{}
	var sDTO *services.AccountDTO
	// 开启一个事务
	err := base.DB().Transaction(func(tx *gorm.DB) error {
		accountDao.db = tx
		accountLogDao.db = tx
		// 插入账户数据
		id, err := accountDao.Insert(&d.account)
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户失败")
		}
		// 账户创建成功 则创建账户流水
		id, err = accountLogDao.Insert(&d.accountLog)
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户流水失败")
		}
		account, _ := accountDao.GetOne(d.account.AccountNo)
		d.account = *account
		return nil
	})
	sDTO = d.account.ToDTO()
	return sDTO, err
}

// 转账
func (d *domain) Transfer(dto services.AccountTransferDTO) (status services.TransferedStatus, err error) {
	amount := dto.Amount
	// 如果是支出，金额应该为负数
	if dto.ChangeFlag == services.FlagTransferOut {
		amount = amount.Mul(decimal.NewFromFloat(-1))
	}
	// 创建流水记录
	d.accountLog = Log{}
	d.accountLog.FromTransferDTO(&dto)
	// 生成账户流水编号
	d.createAccountLogNo()
	err = base.DB().Transaction(func(tx *gorm.DB) error {
		accountDao := AccountsDao{db: tx}
		accountLogDao := LogDao{db: tx}
		// 更新交易主体余额
		id, err := accountDao.UpdateBalance(dto.TradeBody.AccountNo, amount)
		if err != nil {
			return err
		}
		// 没有更新成功并且是支出时
		if id <= 0 && dto.ChangeType == services.EnvelopeOutgoing {
			status = services.TransferedStatusSufficientFunds
			return errors.New("账户余额不足")
		}
		account, err := accountDao.GetOne(dto.TradeBody.AccountNo)
		if err != nil {
			return err
		}
		if account.Id <= 0 {
			return errors.New("账户不存在")
		}
		// 设置账户流水的剩余余额
		d.accountLog.Balance = account.Balance
		// 更新余额成功 写入流水记录
		id, err = accountLogDao.Insert(&d.accountLog)
		if err != nil {
			status = services.TransferedStatusFailure
			return errors.New("创建账户流水失败")
		}
		if id <= 0 {
			status = services.TransferedStatusFailure
			return errors.New("创建账户流水失败")
		}
		status = services.TransferedStatusSuccess
		return nil
	})
	return status, err
}

// 根据账户编号查询账户信息
func (d *domain) GetAccountByAccountNo(accountNo string) (*services.AccountDTO, error) {
	dto := &services.AccountDTO{}
	err := base.DB().Transaction(func(tx *gorm.DB) error {
		accountDao := AccountsDao{db: tx}
		account, err := accountDao.GetOne(accountNo)
		if err != nil {
			return err
		}
		if account.Id <= 0 {
			return errors.New("账户不存在")
		}
		dto = account.ToDTO()
		return nil
	})
	return dto, err
}

func (d *domain) GetEnvelopeAccountByUserId(userId string) (*services.AccountDTO, error) {
	dto := &services.AccountDTO{}
	err := base.DB().Transaction(func(tx *gorm.DB) error {
		accountDao := AccountsDao{db: tx}
		account, err := accountDao.GetAccountByUserIdAcoountType(userId, int(services.EnvelopeAccount))
		if err != nil {
			return err
		}
		if account.Id <= 0 {
			return errors.New("账户不存在")
		}
		dto = account.ToDTO()
		return nil
	})
	return dto, err
}
