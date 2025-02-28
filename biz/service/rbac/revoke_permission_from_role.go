package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type RevokePermissionFromRoleService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRevokePermissionFromRoleService(Context context.Context, RequestContext *app.RequestContext) *RevokePermissionFromRoleService {
	return &RevokePermissionFromRoleService{RequestContext: RequestContext, Context: Context}
}

func (h *RevokePermissionFromRoleService) Run(req *rbac.RevokePermissionReq) (resp *rbac.RevokePermissionResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
