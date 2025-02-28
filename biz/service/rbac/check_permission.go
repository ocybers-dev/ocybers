package rbac

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type CheckPermissionService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckPermissionService(Context context.Context, RequestContext *app.RequestContext) *CheckPermissionService {
	return &CheckPermissionService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckPermissionService) Run(req *rbac.CheckPermissionReq) (resp *rbac.CheckPermissionResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
