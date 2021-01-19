package account

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"goEnvelope/infra/base"
	"goEnvelope/services"
	"testing"
)

func TestDomain_Create(t *testing.T) {
	dto := services.AccountDTO{
		AccountCreateDTO: services.AccountCreateDTO{
			UserId: ksuid.New().Next().String(),
			Username: sql.NullString{
				String: "测试用户3",
				Valid:  true,
			},
			Amount: decimal.NewFromFloat(3400),
			Status: 1,
		},
	}
	Convey("测试创建账户", t, func() {
		domain := &domain{}
		result, err := domain.Create(dto)
		So(err, ShouldBeNil)
		So(result.UserId, ShouldEqual, dto.UserId)
		So(result.Amount, ShouldEqual, dto.Amount)
		So(result.Status, ShouldEqual, dto.Status)
	})
}

func TestDomain_Transfer(t *testing.T) {
	accountNo := "1nBfueECblmEcNA7hOP1VEQdHSD"
	targetAccountNo := "1nCXBnM3ZNXx8WDpSbAfrF0DYCo"
	accountDao := AccountsDao{db: base.DB()}
	account, _ := accountDao.GetOne(accountNo)
	targetAccount, _ := accountDao.GetOne(targetAccountNo)
	dto := services.AccountTransferDTO{
		TradeNo: ksuid.New().Next().String(),
		TradeBody: services.TradeParticipator{
			AccountNo: account.AccountNo,
			Username:  account.Username,
			UserId:    account.UserId,
		},
		TradeTarget: services.TradeParticipator{
			AccountNo: targetAccount.AccountNo,
			Username:  targetAccount.Username,
			UserId:    targetAccount.UserId,
		},
		Amount:     decimal.NewFromFloat(10),
		ChangeType: services.EnvelopeOutgoing,
		ChangeFlag: services.FlagTransferOut,
		Desc:       "测试转账-支出",
	}
	Convey("测试转账-余额充足", t, func() {
		domain := &domain{}
		status, err := domain.Transfer(dto)
		balance := account.Balance.Sub(dto.Amount)
		account, _ = accountDao.GetOne(accountNo)
		So(err, ShouldBeNil)
		So(status, ShouldEqual, services.TransferedStatusSuccess)
		So(account.Balance, ShouldEqual, balance)
	})
	Convey("测试转账-余额不足", t, func() {
		domain := &domain{}
		dto.Amount = decimal.NewFromFloat(100)
		status, err := domain.Transfer(dto)
		newAccount, _ := accountDao.GetOne(accountNo)
		So(err, ShouldNotBeNil)
		So(status, ShouldEqual, services.TransferedStatusSufficientFunds)
		So(account.Balance, ShouldEqual, newAccount.Balance)
	})
}

func TestDomain_GetAccountByAccountNo(t *testing.T) {
	Convey("测试根据账户编号查询账户", t, func() {
		accountNo := "1nBfueECblmEcNA7hOP1VEQdHSD"
		domain := &domain{}
		account, err := domain.GetAccountByAccountNo(accountNo)
		So(err, ShouldBeNil)
		t.Logf("account: %+v", account)
	})
}

func TestDomain_GetEnvelopeAccountByUserId(t *testing.T) {
	Convey("测试根据用户Id红包账户", t, func() {
		userId := "1nBfujkBCWrpD4seMIUt4v7A5f2"
		domain := &domain{}
		account, err := domain.GetEnvelopeAccountByUserId(userId)
		So(err, ShouldBeNil)
		t.Logf("account: %+v", account)
	})
}
