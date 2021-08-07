// Package model is to init the database
package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"mbook/internal/pkg/setting"
)

// NewDBEngine 因为这里不属于重点，所以现在直接把配置写死
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, errors.Wrap(err, "Sql Open error")
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	return db, nil
}
