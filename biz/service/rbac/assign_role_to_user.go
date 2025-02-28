package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type AssignRoleToUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAssignRoleToUserService(Context context.Context, RequestContext *app.RequestContext) *AssignRoleToUserService {
	return &AssignRoleToUserService{RequestContext: RequestContext, Context: Context}
}

func (h *AssignRoleToUserService) Run(req *rbac.AssignRoleReq) (resp *rbac.AssignRoleResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
