package rbac

import (
	"context"
	rbac2 "github.com/ocybers-dev/ocybers/biz/service/rbac"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/ocybers-dev/ocybers/biz/utils"
	rbac "github.com/ocybers-dev/ocybers/hertz_gen/ocybers/rbac"
)

// CreateRole .
// @router /rbac/create_role [POST]
func CreateRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.CreateRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.CreateRoleResp{}
	resp, err = rbac2.NewCreateRoleService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteRole .
// @router /rbac/delete_role [POST]
func DeleteRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.DeleteRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.DeleteRoleResp{}
	resp, err = rbac2.NewDeleteRoleService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AssignPermissionToRole .
// @router /rbac/assign_permission [POST]
func AssignPermissionToRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.AssignPermissionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.AssignPermissionResp{}
	resp, err = rbac2.NewAssignPermissionToRoleService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// RevokePermissionFromRole .
// @router /rbac/revoke_permission [POST]
func RevokePermissionFromRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.RevokePermissionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.RevokePermissionResp{}
	resp, err = rbac2.NewRevokePermissionFromRoleService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AssignRoleToUser .
// @router /rbac/assign_role_to_user [POST]
func AssignRoleToUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.AssignRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.AssignRoleResp{}
	resp, err = rbac2.NewAssignRoleToUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// RevokeRoleFromUser .
// @router /rbac/revoke_role_from_user [POST]
func RevokeRoleFromUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.RevokeRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.RevokeRoleResp{}
	resp, err = rbac2.NewRevokeRoleFromUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CheckPermission .
// @router /rbac/check_permission [GET]
func CheckPermission(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.CheckPermissionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.CheckPermissionResp{}
	resp, err = rbac2.NewCheckPermissionService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetRolePermissions .
// @router /rbac/get_role_permissions [GET]
func GetRolePermissions(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.GetRolePermissionsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.GetRolePermissionsResp{}
	resp, err = rbac2.NewGetRolePermissionsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetUserRoles .
// @router /rbac/get_user_roles [GET]
func GetUserRoles(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.GetUserRolesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.GetUserRolesResp{}
	resp, err = rbac2.NewGetUserRolesService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetAllRoles .
// @router /rbac/get_all_roles [GET]
func GetAllRoles(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rbac.GetAllRolesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &rbac.GetAllRolesResp{}
	resp, err = rbac2.NewGetAllRolesService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
