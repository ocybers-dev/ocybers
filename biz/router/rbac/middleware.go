// Code generated by hertz generator.

package rbac

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ocybers-dev/ocybers/biz/dal/casbin"
)

func rootMw() []app.HandlerFunc {
	return nil
}

func _rbacMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		casbin.AutoDBRoleMW(), // 应用Casbin中间件
	}
}

func _assignpermissiontoroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _assignroletouserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkpermissionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getallrolesMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getrolepermissionsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserrolesMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _revokepermissionfromroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _revokerolefromuserMw() []app.HandlerFunc {
	// your code...
	return nil
}
