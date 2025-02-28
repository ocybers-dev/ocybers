package rbac

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type AssignPermissionToRoleService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAssignPermissionToRoleService(Context context.Context, RequestContext *app.RequestContext) *AssignPermissionToRoleService {
	return &AssignPermissionToRoleService{RequestContext: RequestContext, Context: Context}
}

func (h *AssignPermissionToRoleService) Run(req *rbac.AssignPermissionReq) (resp *rbac.AssignPermissionResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
