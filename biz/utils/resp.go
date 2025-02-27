package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	c.JSON(code, struct {
		IsError bool   `json:"error"`
		ErrMeg  string `json:"err_meg"`
	}{
		IsError: true,
		ErrMeg:  err.Error(),
	})
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}
