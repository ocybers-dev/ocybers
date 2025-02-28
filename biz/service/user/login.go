package user

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/duke-git/lancet/v2/random"
	"github.com/hertz-contrib/paseto"
	"github.com/ocybers-dev/ocybers/biz/dal/model"
	"github.com/ocybers-dev/ocybers/biz/dal/mysql"
	"github.com/ocybers-dev/ocybers/biz/dal/query"
	"github.com/ocybers-dev/ocybers/conf"
	user "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "请求参数 = %+v", req)
		hlog.CtxInfof(h.Context, "响应结果 = %+v", resp)
	}()

	u := query.Use(mysql.DB).User
	q := u.WithContext(h.Context)

	// 1. 根据用户名或邮箱查找用户
	var existingUser *model.User
	if req.Username != "" {
		existingUser, err = q.Where(u.Username.Eq(req.Username)).First()
	} else if req.Email != "" {
		existingUser, err = q.Where(u.Email.Eq(req.Email)).First()
	} else if req.Phone != "" {
		existingUser, err = q.Where(u.Phone.Eq(req.Phone)).First()
	} else {
		hlog.CtxInfof(h.Context, "无登录账号")
		return nil, errors.New("登录方式不能为空，请输入用户名、邮箱或手机号")
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hlog.CtxInfof(h.Context, "用户不存在")
			return nil, errors.New("用户不存在")
		}
		hlog.CtxErrorf(h.Context, "查找用户失败: %v", err)
		return nil, errors.New("内部服务器错误")
	}

	// 2. 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(req.Password))
	if err != nil {
		hlog.CtxInfof(h.Context, "密码错误")
		return nil, errors.New("密码错误")
	}

	// 3. 生成token
	now := time.Now()
	uuid, _ := random.UUIdV4()
	genTokenFunc, err := paseto.NewV4EncryptFunc(conf.GetConf().Hertz.PaseToSymmetricKey, []byte(conf.GetConf().Hertz.PaseToImplicit))
	if err != nil {
		hlog.CtxErrorf(h.Context, "生成token函数失败: %v", err)
		return nil, errors.New("内部服务器错误1")
	}
	token, err := genTokenFunc(&paseto.StandardClaims{
		Issuer:    conf.GetConf().Hertz.PaseToIssuer,
		Subject:   existingUser.ID,
		Audience:  "Ocybers端",
		Jti:       uuid,
		ExpiredAt: now.Add(time.Duration(conf.GetConf().Hertz.PaseToExpired) * time.Hour),
		NotBefore: now,
		IssuedAt:  now,
	}, nil, nil)
	// token, err := paseto.DefaultGenTokenFunc()(&paseto.StandardClaims{
	// 	Issuer:    conf.GetConf().Hertz.PaseToIssuer,
	// 	Subject:   existingUser.ID,
	// 	Audience:  "Ocybers端",
	// 	Jti:       uuid,
	// 	ExpiredAt: now.Add(time.Duration(conf.GetConf().Hertz.PaseToExpired) * time.Hour),
	// 	NotBefore: now,
	// 	IssuedAt:  now,
	// }, nil, nil)
	// if err != nil {
	// 	hlog.Error("生成token失败: %v", err)
	// 	return nil, errors.New("内部服务器错误")
	// }

	// 4. 更新最后登录时间
	existingUser.LastLoginAt = &now
	err = q.Save(existingUser)
	if err != nil {
		hlog.CtxErrorf(h.Context, "更新最后登录时间失败: %v", err)
		return nil, errors.New("内部服务器错误")
	}

	// 5. 构造响应
	resp = &user.LoginResp{
		UserId: existingUser.ID,
		Token:  token,
	}

	hlog.CtxInfof(h.Context, "用户登录成功，用户ID: %s", existingUser.ID)
	return resp, nil
}
