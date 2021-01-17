package account

import "gorm.io/gorm"

type LogDao struct {
	db *gorm.DB
}

// 通过流水编号查询流水记录
func (dao *LogDao) GetOne(logNo string) (l *Log, err error) {
	var log Log
	err = dao.db.Where("log_no = ?", logNo).Find(&log).Error
	return &log, err
}

// 通过交易编号
func (dao *LogDao) GetOneByTradeNo(tradeNo string) (l *Log, err error) {
	var log Log
	err = dao.db.Where("trade_no = ?", tradeNo).Find(&log).Error
	return &log, err
}

// 新增流水记录
func (dao *LogDao) Insert(l *Log) (id int64, err error) {
	tx := dao.db.Create(l)
	return tx.RowsAffected, tx.Error
}
