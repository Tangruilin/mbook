package global

import (
	"github.com/jinzhu/gorm"
	"mbook/internal/pkg/setting"
)

var (
	DBEngine        *gorm.DB
	DatabaseSetting = &setting.DatabaseSettingS{
		DBType:       "mysql",
		UserName:     "root",
		Password:     "Tangruilin123",
		Host:         "127.0.0.1:3306",
		DBName:       "mbook",
		Charset:      "utf8",
		ParseTime:    true,
		MaxIdleConns: 10,
		MaxOpenConns: 30,
	}
)
