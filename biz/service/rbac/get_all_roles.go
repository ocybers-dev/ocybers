package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

type GetAllRolesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAllRolesService(Context context.Context, RequestContext *app.RequestContext) *GetAllRolesService {
	return &GetAllRolesService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAllRolesService) Run(req *rbac.GetAllRolesReq) (resp *rbac.GetAllRolesResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
