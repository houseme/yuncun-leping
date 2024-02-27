package authorize

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/yuncun-leping/app/front/api/authorize/v1"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
)

// Authorize is the handler for Authorize Controller action.
func (c *ControllerV1) Authorize(ctx context.Context, req *v1.AuthorizeReq) (res *v1.AuthorizeRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-authorize-Authorize")
	defer span.End()

	res = &v1.AuthorizeRes{}
	res.AuthorizeOutput, err = service.Authorize().Authorization(ctx, req.AuthorizeInput)
	return
}
