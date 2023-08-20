package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/houseme/yuncun-leping/app/front/api/hello/v1"
)

// Controller struct
type Controller struct{}

// New Controller
func New() *Controller {
	return &Controller{}
}

// Hello say hello
func (c *Controller) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
