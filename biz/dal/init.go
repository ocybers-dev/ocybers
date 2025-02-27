package dal

import (
	"github.com/ocybers-dev/ocybers/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
