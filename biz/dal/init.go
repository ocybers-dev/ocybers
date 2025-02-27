package dal

import (
	"github.com/ocybers-dev/ocybers/biz/dal/mysql"
	"github.com/ocybers-dev/ocybers/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
