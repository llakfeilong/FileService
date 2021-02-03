package dbManger

import (
	"github.com/jinzhu/gorm"
)

//创建
func CreateStruct(value interface{}) error {
	tx := db.Begin()
	if err := tx.Create(value).Error; err != nil {
		tx.Rollback() //回滚事务
		return err
	}
	tx.Commit() //提交事务
	return nil
}

//更新
func UpdateStruct(value interface{}) error {
	tx := db.Begin()
	if err := tx.Save(value).Error; err != nil {
		tx.Rollback() //回滚事务
		return err
	}
	tx.Commit() //提交事务
	return nil
}

//删除
func DeleteStruct(value interface{}) error {
	tx := db.Begin()
	if err := tx.Delete(value).Error; err != nil {
		tx.Rollback() //回滚事务
		return err
	}
	tx.Commit() //提交事务
	return nil
}

//func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
//	return db.Where("amount > ?", 1000)
//}
//
//func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
//	return db.Where("pay_mode_sign = ?", "C")
//}

//分页查询
func PageQuery(scopes []func(*gorm.DB) *gorm.DB, pageSize, pageNo int, order string) *gorm.DB {
	return db.Scopes(scopes...).Limit(pageSize).Offset((pageNo - 1) * pageSize).Order(order)
}
