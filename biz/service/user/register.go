package user

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/duke-git/lancet/v2/random"
	"github.com/ocybers-dev/ocybers/biz/dal/model"
	"github.com/ocybers-dev/ocybers/biz/dal/mysql"
	"github.com/ocybers-dev/ocybers/biz/dal/query"
	user "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "请求参数 = %+v", req)
		hlog.CtxInfof(h.Context, "响应结果 = %+v", resp)
	}()

	u := query.Use(mysql.DB).User
	q := u.WithContext(h.Context)

	// 1. 检查用户名是否已存在
	existingUser, err := q.Where(u.Username.Eq(req.Username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		hlog.CtxErrorf(h.Context, "检查用户名是否存在失败: %v", err)
		return nil, errors.New("内部服务器错误")
	}
	if existingUser != nil {
		hlog.CtxInfof(h.Context, "用户名 %s 已被占用", req.Username)
		return nil, errors.New("用户名已被占用")
	}

	// 2. 检查邮箱是否已存在
	if req.Email == "" {
		hlog.CtxWarnf(h.Context, "邮箱为空，跳过邮箱唯一性检查")
	} else {
		existingUser, err = q.Where(u.Email.Eq(req.Email)).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			hlog.CtxErrorf(h.Context, "检查邮箱是否存在失败: %v", err)
			return nil, errors.New("内部服务器错误")
		}
		if existingUser != nil {
			hlog.CtxInfof(h.Context, "邮箱 %s 已被占用", req.Email)
			return nil, errors.New("邮箱已被占用")
		}
	}

	// 3. 对密码进行哈希
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		hlog.CtxErrorf(h.Context, "密码哈希失败: %v", err)
		return nil, errors.New("内部服务器错误")
	}

	// 4. 生成 UUID
	uuid, err := random.UUIdV4()
	if err != nil {
		hlog.CtxErrorf(h.Context, "生成 UUID 失败: %v", err)
		return nil, errors.New("内部服务器错误")
	}

	// 5. 创建新用户
	newUser := &model.User{
		ID:           string(uuid), // 生成 UUID
		Username:     req.Username,
		Email:        &req.Email, // 指针类型，需取地址
		PasswordHash: string(passwordHash),
		FullName:     nil, // 可选字段，初始为 nil
		Phone:        nil, // 可选字段，初始为 nil
		Status:       1,   // 1=active
		CreatedAt:    time.Now(),
		UpdatedAt:    nil, // 初始为 nil
		LastLoginAt:  nil, // 初始为 nil
	}

	err = q.Create(newUser)
	if err != nil {
		hlog.CtxErrorf(h.Context, "创建用户失败: %v", err)
		return nil, errors.New("内部服务器错误")
	}

	// 6. 构造响应
	resp = &user.RegisterResp{
		UserId: newUser.ID,
	}
	hlog.CtxInfof(h.Context, "用户注册成功，用户ID: %s", newUser.ID)
	return resp, nil
}
