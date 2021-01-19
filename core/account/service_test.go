package account

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"goEnvelope/services"
	"testing"
)

func TestService_CreateAccount(t *testing.T) {
	s := new(service)
	dto := services.AccountCreateDTO{
		UserId:       ksuid.New().Next().String(),
		Username:     sql.NullString{String: "王武", Valid: true},
		AccountName:  "中国银行储蓄卡",
		AccountType:  services.EnvelopeAccount,
		CurrencyCode: "CNY",
		Amount:       decimal.NewFromFloat(2000),
	}
	Convey("测试service:创建账户", t, func() {
		account, err := s.CreateAccount(dto)
		t.Log(account)
		So(err, ShouldBeNil)
	})
}

func TestService_Transfer(t *testing.T) {
	s := &service{}
	account, _ := s.GetAccount("1nI1m5xynxk8pHGvOCPczPipSDP")
	targetAccount, _ := s.GetAccount("1nCXBnM3ZNXx8WDpSbAfrF0DYCo")
	dto := services.AccountTransferDTO{
		TradeNo:     ksuid.New().Next().String(),
		TradeBody:   services.TradeParticipator{
			AccountNo: account.AccountNo,
			UserId:    account.UserId,
			Username:  account.Username,
		},
		TradeTarget: services.TradeParticipator{
			AccountNo: targetAccount.AccountNo,
			UserId:    targetAccount.UserId,
			Username:  targetAccount.Username,
		},
		Amount:      decimal.NewFromFloat(10),
		ChangeType:  services.EnvelopeOutgoing,
		ChangeFlag:  services.FlagTransferOut,
		Desc:        "测试service转账",
	}
	Convey("测试service:转账", t, func() {
		status, err := s.Transfer(dto)
		So(err, ShouldBeNil)
		So(status, ShouldEqual, services.TransferedStatusSuccess)
	})
}

func TestService_GetEnvelopeAccountByUserId(t *testing.T) {
	s := &service{}
	Convey("测试service：查询红包账户信息", t, func() {
		account, _ := s.GetAccount("1nI1m5xynxk8pHGvOCPczPipSDP")
		dto, err := s.GetEnvelopeAccountByUserId(account.UserId)
		So(err, ShouldBeNil)
		t.Log(dto)
	})
}
