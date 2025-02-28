package rbac

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type GetUserRolesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserRolesService(Context context.Context, RequestContext *app.RequestContext) *GetUserRolesService {
	return &GetUserRolesService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserRolesService) Run(req *rbac.GetUserRolesReq) (resp *rbac.GetUserRolesResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
