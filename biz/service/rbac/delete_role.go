package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type DeleteRoleService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteRoleService(Context context.Context, RequestContext *app.RequestContext) *DeleteRoleService {
	return &DeleteRoleService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteRoleService) Run(req *rbac.DeleteRoleReq) (resp *rbac.DeleteRoleResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
