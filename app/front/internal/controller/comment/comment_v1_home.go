package comment

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/yuncun-leping/app/front/api/comment/v1"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
)

// Home is the handler for Comment Controller action.
func (c *ControllerV1) Home(ctx context.Context, req *v1.HomeReq) (res *v1.HomeRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-comment-home")
	defer span.End()

	res = &v1.HomeRes{}
	res.HomeOutput, err = service.Comment().Home(ctx, req.HomeInput)
	return
}
