package account

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"goEnvelope/infra/base"
	"goEnvelope/services"
	"testing"
)

func TestLogDao_Insert(t *testing.T) {
	accountsDao := &AccountsDao{db: base.DB()}
	logDao := LogDao{db: base.DB()}
	Convey("测试新增账户流水", t, func() {
		account, _ := accountsDao.GetOne("1nBdcTYhQJ3oo8YX2KBlVTaOLIC")
		targetAccount, _ := accountsDao.GetOne("1nBfql9f8RwNcw3F6s3gNNvp02X")
		amount := decimal.NewFromFloat(20)
		l := &Log{
			TradeNo:         ksuid.New().Next().String(),
			LogNo:           ksuid.New().Next().String(),
			AccountNo:       account.AccountNo,
			UserId:          account.UserId,
			Username:        account.Username,
			TargetAccountNo: targetAccount.AccountNo,
			TargetUsername:  targetAccount.Username,
			TargetUserId:    targetAccount.UserId,
			Amount:          amount,
			Balance:         account.Balance.Sub(amount),
			ChangeType:      services.EnvelopeOutgoing,
			ChangeFlag:      services.FlagTransferOut,
			Status:          1,
			Desc:            "微信转账",
		}
		id, err := logDao.Insert(l)
		So(id, ShouldBeGreaterThan, 0)
		So(err, ShouldBeNil)
	})
}

func TestLogDao_GetOne(t *testing.T) {
	dao := LogDao{db: base.DB()}
	Convey("测试根据流水编号查询账户流水", t, func() {
		log, err := dao.GetOne("1nCLQQsGQd9MLbOLOClIa5kseeM")
		So(err, ShouldBeNil)
		t.Logf("accountLog: %+v \n", log)
	})
}

func TestLogDao_GetOneByTradeNo(t *testing.T) {
	dao := LogDao{db: base.DB()}
	Convey("测试根据交易编号查询账户流水", t, func() {
		log, err := dao.GetOneByTradeNo("1nCLQUUop8mtnWpKrtBeKp239eU")
		So(err, ShouldBeNil)
		t.Logf("accountLog: %+v \n", log)
	})
}
