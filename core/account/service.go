package account

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"goEnvelope/infra/base"
	"goEnvelope/services"
	"sync"
)

type service struct{}

// 这里没有实际作用只是为了快速实现AccountService接口的方法
var _ services.AccountService = new(service)
var once sync.Once
func init()  {
	// 只被实例化一次
	once.Do(func() {
		services.IAccountService = new(service)
	})
}

// 创建账户
func (s *service) CreateAccount(dto services.AccountCreateDTO) (*services.AccountDTO, error) {
	domain := &domain{}
	// 验证参数
	err := base.Validate().Struct(&dto)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			logrus.Errorf("验证错误:%v \n", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				logrus.Info(err.Translate(base.Translator()))
			}
		}
		return nil, err
	}
	account := services.AccountDTO{
		AccountCreateDTO: services.AccountCreateDTO{
			UserId:       dto.UserId,
			Username:     dto.Username,
			AccountName:  dto.AccountName,
			AccountType:  dto.AccountType,
			CurrencyCode: dto.CurrencyCode,
			Status:       1,
			Amount:       dto.Amount,
		},
	}
	result, err := domain.Create(account)
	return result, err
}

// 转账
func (s *service) Transfer(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	domain := &domain{}
	// 验证参数
	err := base.Validate().Struct(&dto)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			logrus.Errorf("验证错误:%v \n", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				logrus.Info(err.Translate(base.Translator()))
			}
		}
		return services.TransferedStatusFailure, err
	}
	// 支出 转账类型 必须<= -1
	if dto.ChangeFlag == services.FlagTransferOut {
		// 参数有误
		if dto.ChangeType > 0 {
			return services.TransferedStatusFailure, errors.New("资金变化为ChangeFlag支出时，转账类型ChangeType必须<=-1")
		}
	} else if dto.ChangeFlag == services.FLagTransferIn {
		if dto.ChangeType < 0 {
			return services.TransferedStatusFailure, errors.New("资金变化为ChangeFlag收入时，转账类型ChangeType必须>0")
		}
	}
	return domain.Transfer(dto)
}

// 储值
func (s *service) StoreValue(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	dto.TradeTarget = dto.TradeBody
	dto.ChangeType = services.AccountStoreValue
	dto.ChangeFlag = services.FLagTransferIn
	return s.Transfer(dto)
}

// 根据账户编号 获取账户信息
func (s *service) GetAccount(accountNo string) (*services.AccountDTO, error) {
	domain := &domain{}
	return domain.GetAccountByAccountNo(accountNo)
}

// 根据用户id查询红包账户信息
func (s *service) GetEnvelopeAccountByUserId(userId string) (*services.AccountDTO, error) {
	domain := &domain{}
	return domain.GetEnvelopeAccountByUserId(userId)
}
