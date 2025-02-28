package casbin

import (
	"context"
	"fmt"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
	_casbin "github.com/hertz-contrib/casbin"
	"github.com/ocybers-dev/ocybers/conf"
)

var (
	CasbinMiddleware *_casbin.Middleware
	E                *casbin.Enforcer
)

// Init 初始化Casbin中间件
func Init() {
	// 获取数据库连接字符串（DSN）
	dsn := conf.GetConf().MySQL.DSN

	// 初始化Gorm适配器
	adapter, err := initGormAdapter(dsn)
	if err != nil {
		hlog.Fatalf("初始化Gorm适配器失败: %v", err)
		return
	}

	// 初始化Casbin模型
	model, err := initCasbinModel()
	if err != nil {
		hlog.Fatalf("加载Casbin模型失败: %v", err)
		return
	}

	// 创建Casbin Enforcer
	E, err := casbin.NewEnforcer(model, adapter)
	if err != nil {
		hlog.Fatalf("初始化Casbin Enforcer失败: %v", err)
		return
	}

	// 初始化Casbin中间件
	CasbinMiddleware, err = _casbin.NewCasbinMiddlewareFromEnforcer(E, subjectFromContext)
	if err != nil {
		hlog.Fatalf("初始化Casbin中间件失败: %v", err)
	}
}

// initGormAdapter 初始化并返回Gorm适配器
func initGormAdapter(dsn string) (*gormadapter.Adapter, error) {
	// 查找 '/' 的索引位置
	slashIndex := strings.Index(dsn, "/")
	if slashIndex == -1 {
		return nil, fmt.Errorf("无效的数据库连接字符串，未找到 '/' 位置: %v", dsn)
	}

	// 获取连接字符串部分（即协议、主机和端口部分）
	connStrPart := dsn[:slashIndex+1]

	// 获取数据库名称部分（"/" 后面的部分），并去掉查询参数部分（如果有）
	dbName := dsn[slashIndex+1:]
	if queryIndex := strings.Index(dbName, "?"); queryIndex != -1 {
		dbName = dbName[:queryIndex]
	}

	// 创建Gorm适配器
	adapter, err := gormadapter.NewAdapter("mysql", connStrPart, dbName)
	if err != nil {
		return nil, fmt.Errorf("创建Gorm适配器失败: %v", err)
	}

	return adapter, nil
}

// initCasbinModel 初始化并返回Casbin的RBAC模型
func initCasbinModel() (model.Model, error) {
	// 创建模型
	m := model.NewModel()

	// 添加请求定义（r）: 请求包含 3 个元素，sub、obj 和 act
	m.AddDef("r", "r", "sub, obj, act")

	// 添加策略定义（p）: 策略包含 3 个元素，sub、obj 和 act
	m.AddDef("p", "p", "sub, obj, act")

	// 添加角色定义（g）: 用户与角色的关系，角色可以分配给用户
	m.AddDef("g", "g", "_, _")

	// 添加效果定义（e）: 规定策略的生效条件
	m.AddDef("e", "e", "some(where (p.eft == allow))")

	// 添加匹配器定义（m）: 用于匹配规则，判断请求是否符合授权的条件
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	return m, nil
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
