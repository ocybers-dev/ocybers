package casbin

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/duke-git/lancet/v2/convertor"
	_ "github.com/go-sql-driver/mysql"
	_casbin "github.com/hertz-contrib/casbin"
	"github.com/ocybers-dev/ocybers/conf"
	"strings"
)

var (
	middleware *_casbin.Middleware
	E          *casbin.Enforcer
)

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
	casbinModel, err := initCasbinModel()
	if err != nil {
		hlog.Fatalf("加载Casbin模型失败: %v", err)
		return
	}

	// 创建Casbin Enforcer
	E, err = casbin.NewEnforcer(casbinModel, adapter)
	if err != nil {
		hlog.Fatalf("初始化Casbin Enforcer失败: %v", err)
		return
	}

	// 初始化权限数据
	initPermissions()

	// 初始化Casbin中间件
	middleware, err = _casbin.NewCasbinMiddlewareFromEnforcer(E, subjectFromContext)
	if err != nil {
		hlog.Fatalf("初始化Casbin中间件失败: %v", err)
	}
}

func AutoDBRoleMW() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		path := convertor.ToString(c.Request.Path())
		// 获取匹配当前路径的所有权限策略
		filteredPolicy, err := E.GetFilteredPolicy(1, path)
		if err != nil {
			// 如果获取策略失败，记录日志并返回500错误
			hlog.CtxErrorf(ctx, "获取路径权限失败: %v", err)
			c.JSON(consts.StatusInternalServerError, utils.H{"error": "服务器错误，无法获取权限"})
			c.Abort()
			return
		}

		// 如果没有找到任何角色与路径匹配，返回未授权错误
		if len(filteredPolicy) == 0 {
			hlog.CtxWarnf(ctx, "当前接口无角色开放，请联系开发者: %s", path)
			c.JSON(consts.StatusUnauthorized, utils.H{"error": "当前接口无角色开放，请联系开发者"})
			c.Abort()
			return
		}

		// 获取角色列表
		var roles []string
		for _, policy := range filteredPolicy {
			roles = append(roles, policy[0])
		}
		rolesStr := strings.Join(roles, " ")

		// 调用 Casbin 中间件，进行角色验证
		middleware.RequiresRoles(rolesStr,
			_casbin.WithLogic(_casbin.OR),
			// 如果未通过认证，返回401 Unauthorized
			_casbin.WithUnauthorized(func(ctx context.Context, c *app.RequestContext) {
				hlog.CtxWarnf(ctx, "用户未授权访问接口: %s", path)
				c.JSON(consts.StatusUnauthorized, utils.H{"error": "未授权"})
				c.Abort()
			}),
			// 如果认证通过，但权限不足，返回403 Forbidden
			_casbin.WithForbidden(func(ctx context.Context, c *app.RequestContext) {
				hlog.CtxWarnf(ctx, "用户权限不足，无法访问接口: %s", path)
				c.JSON(consts.StatusForbidden, utils.H{"error": "权限不足"})
				c.Abort()
			}),
		)(ctx, c)
	}
}

// initPermissions 初始化路径权限
func initPermissions() {
	// 预设的权限列表，包含路径和方法
	defaultPolicies := [][]string{
		// 用户访问权限
		{"admin", "/ping", "GET"},
		{"user", "/ping", "GET"},

		// RBAC服务接口权限
		{"admin", "/rbac/create_role", "POST"},
		{"admin", "/rbac/delete_role", "POST"},
		{"admin", "/rbac/assign_permission", "POST"},
		{"admin", "/rbac/revoke_permission", "POST"},
		{"admin", "/rbac/assign_role_to_user", "POST"},
		{"admin", "/rbac/revoke_role_from_user", "POST"},

		// 权限检查接口
		{"user", "/rbac/check_permission", "GET"},
		{"admin", "/rbac/check_permission", "GET"},

		// 获取角色相关接口
		{"admin", "/rbac/get_role_permissions", "GET"},
		{"admin", "/rbac/get_user_roles", "GET"},
		{"admin", "/rbac/get_all_roles", "GET"},

		// 根据需求添加更多的权限配置...
	}

	for _, policy := range defaultPolicies {
		// 检查权限是否已存在
		exists, err := E.Enforce(policy[0], policy[1], policy[2])
		if err != nil {
			hlog.Errorf("检查权限是否存在时发生错误: %v", err)
			continue
		}

		if exists {
			// 如果权限已存在，跳过
			continue
		}

		// 如果权限不存在，插入到数据库
		_, err = E.AddPolicy(policy[0], policy[1], policy[2])
		if err != nil {
			hlog.Errorf("插入权限到数据库失败: %v", err)
		} else {
			hlog.Infof("成功插入权限: %v", policy)
		}
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
