package web

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"goEnvelope/infra"
	"goEnvelope/infra/base"
	"goEnvelope/services"
)

func init() {
	infra.RegisterApi(&AccountApi{})
}

type AccountApi struct {}

func (a *AccountApi) Init()  {
	// 定义统一的前缀
	groupRouter := base.Iris().Party("/v1/account")
	groupRouter.Post("/create", createHandler)
	groupRouter.Post("/transfer", transferHandler)
	groupRouter.Get("/{accountNo}", getAccount)
	groupRouter.Get("/envelope/{userId}", getEnvelopeAccount)
}

// 账户创建
func createHandler(ctx iris.Context) {
	account := services.AccountCreateDTO{}
	// 获取参数
	err := ctx.ReadJSON(&account)
	r := base.Response{
		Code:    base.ResponseCodeOk,
		Message: "success",
		Data:    nil,
	}
	if err != nil {
		r.Code = base.ResponseCodeRequestParamsError
		r.Message = err.Error()
		_, _ = ctx.JSON(&r)
		return
	}
	// 创建账户
	service := services.GetAccountService()
	dto, err := service.CreateAccount(account)
	if err != nil {
		r.Code = base.ResponseCodeRequestParamsError
		r.Message = err.Error()
		_, _ = ctx.JSON(&r)
		return
	}
	r.Data = dto
	_, _ = ctx.JSON(&r)
}

// 转账
func transferHandler(ctx iris.Context)  {
	transfer := services.AccountTransferDTO{}
	err := ctx.ReadJSON(&transfer)
	r := base.Response{
		Code:    base.ResponseCodeOk,
		Message: "success",
		Data:    nil,
	}
	if err != nil {
		r.Code = base.ResponseCodeRequestParamsError
		r.Message = "参数错误"
		_, _ = ctx.JSON(r)
		return
	}
	service := services.GetAccountService()
	status, err := service.Transfer(transfer)
	if err != nil {
		r.Code = base.ResponseCodeInterServerError
		r.Message = err.Error()
		_, _ = ctx.JSON(r)
		return
	}
	if status != services.TransferedStatusSuccess {
		// 业务异常
		r.Code = base.ResponseCodeBizError
	}
	r.Data = status
	_, _ = ctx.JSON(&r)
}

// 查询红包账户
func getEnvelopeAccount(ctx iris.Context)  {
	userId := ctx.Params().Get("userId")
	r := base.Response{
		Code:    base.ResponseCodeOk,
		Message: "success",
		Data:    nil,
	}
	if userId == "" {
		r.Code = base.ResponseCodeRequestParamsError
		r.Message = "未知用户编号"
		_, _ = ctx.JSON(&r)
		return
	}
	service := services.GetAccountService()
	account, err := service.GetEnvelopeAccountByUserId(userId)
	if err != nil {
		r.Code = base.ResponseCodeBizError
		r.Message = err.Error()
		_, _ = ctx.JSON(&r)
		return
	}
	r.Data = *account
	_, _ = ctx.JSON(&r)
}

// 查询账户
func getAccount(ctx iris.Context)  {
	accountNo := ctx.Params().Get("accountNo")
	fmt.Println(accountNo)
	r := base.Response{
		Code:    base.ResponseCodeOk,
		Message: "success",
		Data:    nil,
	}
	if accountNo == "" {
		r.Code = base.ResponseCodeRequestParamsError
		r.Message = "未知账户编号"
		_, _ = ctx.JSON(&r)
		return
	}
	service := services.GetAccountService()
	accountDTO, err := service.GetAccount(accountNo)
	if err != nil {
		r.Code = base.ResponseCodeBizError
		r.Message = "未查询到账户"
		_, _ = ctx.JSON(&r)
		return
	}
	r.Data = *accountDTO
	_, _ = ctx.JSON(&r)
}
