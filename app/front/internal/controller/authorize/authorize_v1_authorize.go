package authorize

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/houseme/yuncun-leping/app/front/api/authorize/v1"
)

// Authorize is the handler for Authorize Controller action.
func (c *ControllerV1) Authorize(ctx context.Context, req *v1.AuthorizeReq) (res *v1.AuthorizeRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
