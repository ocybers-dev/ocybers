package rbac

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type GetRolePermissionsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetRolePermissionsService(Context context.Context, RequestContext *app.RequestContext) *GetRolePermissionsService {
	return &GetRolePermissionsService{RequestContext: RequestContext, Context: Context}
}

func (h *GetRolePermissionsService) Run(req *rbac.GetRolePermissionsReq) (resp *rbac.GetRolePermissionsResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
