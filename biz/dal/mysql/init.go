package mysql

import (
	"github.com/ocybers-dev/ocybers/biz/dal/model"
	"github.com/ocybers-dev/ocybers/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	// 自动迁移
	err = DB.AutoMigrate(
		&model.User{},
		&model.Article{},
	)
	if err != nil {
		panic(err)
	}
}
