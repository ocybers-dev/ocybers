package casbin

import (
	"context"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hertz-contrib/casbin"
)

var (
	CasbinMiddleware *casbin.Middleware
)

func Init() {
	var err error

	a, _ := gormadapter.NewAdapter("mysql", "gorm:gorm@tcp(127.0.0.1:3306)/", "gorm")

	// 初始化Casbin中间件
	CasbinMiddleware, err = casbin.NewCasbinMiddleware("model.conf", a, subjectFromContext)
	if err != nil {
		hlog.Fatalf("初始化Casbin中间件失败: %v", err)
	}
}

// subjectFromContext 从请求上下文中提取用户标识
func subjectFromContext(ctx context.Context, c *app.RequestContext) string {
	// 获取访问实体
	casbinSubject, exists := c.Get("sub")
	if !exists {
		hlog.CtxWarnf(ctx, "上下文中未找到 'sub' 字段")
		return ""
	}

	// 类型断言，确保casbinSubject是字符串
	subject, ok := casbinSubject.(string)
	if !ok {
		hlog.CtxWarnf(ctx, "'sub' 字段不是字符串类型")
		return ""
	}

	return subject
}
