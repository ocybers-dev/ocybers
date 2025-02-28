package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type RevokeRoleFromUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRevokeRoleFromUserService(Context context.Context, RequestContext *app.RequestContext) *RevokeRoleFromUserService {
	return &RevokeRoleFromUserService{RequestContext: RequestContext, Context: Context}
}

func (h *RevokeRoleFromUserService) Run(req *rbac.RevokeRoleReq) (resp *rbac.RevokeRoleResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
