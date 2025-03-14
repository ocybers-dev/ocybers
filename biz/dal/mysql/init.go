package mysql

import (
	"fmt"

	"github.com/ocybers-dev/ocybers/biz/dal/casbin"
	"github.com/ocybers-dev/ocybers/biz/dal/model"
	"github.com/ocybers-dev/ocybers/conf"
	"golang.org/x/crypto/bcrypt"
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
	)
	if err != nil {
		panic(err)
	}
}

func InitAdminUser() {
	//初始化管理员用户
	// 检查用户表是否为空，如果为空则初始化管理员用户
	var count int64
	DB.Model(&model.User{}).Count(&count)

	if count == 0 {
		adminPassword := "Y2hlbjA0MTY="

		// 创建管理员用户
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		if err != nil {
			panic("生成密码哈希失败: " + err.Error())
		}

		adminUser := model.User{
			ID:           "admin",
			Username:     "ocybers",
			PasswordHash: string(passwordHash),
		}

		result := DB.Create(&adminUser)
		if result.Error != nil {
			panic("初始化管理员用户失败: " + result.Error.Error())
		}

		casbin.E.AddRoleForUser(adminUser.ID, "user")
		casbin.E.AddRoleForUser(adminUser.ID, "admin")

		fmt.Println("已创建管理员用户，用户名: ocybers, 初始密码: " + adminPassword)
		fmt.Println("请在首次登录后立即修改初始密码！")
	}
}
