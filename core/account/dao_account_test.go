package account

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"goEnvelope/infra/base"
	_ "goEnvelope/testx"
	"testing"
)

func TestAccountDao_GetOne(t *testing.T) {
	dao := &AccountsDao{db: base.DB()}
	Convey("查询账户", t, func() {
		accountNo := "1nBdcTYhQJ3oo8YX2KBlVTaOLIC1"
		account, err := dao.GetOne(accountNo)
		t.Logf("account: %+v \n", account)
		t.Logf("err: %v \n", err)
		So(err, ShouldBeNil)
		So(account, ShouldNotBeNil)
	})
}

func TestDao_Insert(t *testing.T) {
	dao := &AccountsDao{db: base.DB()}
	Convey("测试添加账户", t, func() {
		a := &Account{
			Balance:     decimal.NewFromFloat(100),
			Status:      1,
			AccountNo:   ksuid.New().Next().String(),
			AccountName: "测试资金账户",
			UserId:      ksuid.New().Next().String(),
			Username:    sql.NullString{String: "测试用户", Valid: true},
		}
		id, err := dao.Insert(a)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		account, _ := dao.GetOne(a.AccountNo)
		So(account.CreatedAt, ShouldNotBeNil)
		t.Logf("account: %+v ", account)
	})
}

func TestAccountsDao_GetAccountByUserIdAcoountType(t *testing.T) {
	Convey("通过用户编号和账户类型查询账户", t, func() {
		dao := AccountsDao{db: base.DB()}
		a := &Account{
			Balance:     decimal.NewFromFloat(140),
			Status:      2,
			AccountNo:   ksuid.New().Next().String(),
			AccountName: "测试资金账户2",
			UserId:      ksuid.New().Next().String(),
			Username:    sql.NullString{String: "测试用户2", Valid: true},
			AccountType: 2,
		}
		id, err := dao.Insert(a)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		ac, _ := dao.GetAccountByUserIdAcoountType(a.UserId, a.AccountType)
		So(ac.CreatedAt, ShouldNotBeNil)
	})
}

func TestAccountsDao_UpdateBalance(t *testing.T) {
	Convey("测试更新账户余额", t, func() {
		dao := AccountsDao{db: base.DB()}
		accountNo := "1nBdcTYhQJ3oo8YX2KBlVTaOLIC"
		Convey("更新增加余额", func() {
			balance := decimal.NewFromFloat(100.5)
			a := &Account{AccountNo: accountNo, Balance: balance}
			oldRow, _ := dao.GetOne(accountNo)
			id, err := dao.UpdateBalance(a.AccountNo, a.Balance)
			row, _ := dao.GetOne(accountNo)
			So(err, ShouldBeNil)
			So(id, ShouldEqual, 1)
			newBalance := balance.Add(oldRow.Balance)
			So(newBalance, ShouldEqual, row.Balance)
			So(oldRow.CreatedAt, ShouldEqual, row.CreatedAt)
		})
		Convey("更新减少余额，余额充足", func() {
			balance := decimal.NewFromFloat(-20)
			a := &Account{AccountNo: accountNo, Balance: balance}
			oldRow, _ := dao.GetOne(accountNo)
			id, err := dao.UpdateBalance(a.AccountNo, a.Balance)
			row, _ := dao.GetOne(accountNo)
			So(err, ShouldBeNil)
			So(id, ShouldEqual, 1)
			newBalance := oldRow.Balance.Add(balance)
			So(newBalance, ShouldEqual, row.Balance)
			So(oldRow.CreatedAt, ShouldEqual, row.CreatedAt)
		})
		Convey("更新减少余额，余额不足", func() {
			balance := decimal.NewFromFloat(-500)
			a := &Account{AccountNo: accountNo, Balance: balance}
			oldRow, _ := dao.GetOne(accountNo)
			id, err := dao.UpdateBalance(a.AccountNo, a.Balance)
			row, _ := dao.GetOne(accountNo)
			So(err, ShouldNotBeNil)
			So(id, ShouldEqual, 0)
			newBalance := oldRow.Balance.Add(balance)
			So(newBalance, ShouldNotEqual, row.Balance)
		})
	})
}
