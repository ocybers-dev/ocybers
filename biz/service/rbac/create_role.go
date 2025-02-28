package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type CreateRoleService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateRoleService(Context context.Context, RequestContext *app.RequestContext) *CreateRoleService {
	return &CreateRoleService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateRoleService) Run(req *rbac.CreateRoleReq) (resp *rbac.CreateRoleResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
