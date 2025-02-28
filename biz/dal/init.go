package dal

import (
	"github.com/ocybers-dev/ocybers/biz/dal/casbin"
	"github.com/ocybers-dev/ocybers/biz/dal/mysql"
)

func Init() {
	mysql.Init()
	//redis.Init()
	casbin.Init()

}
