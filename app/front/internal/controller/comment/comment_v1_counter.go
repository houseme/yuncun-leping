package comment

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/yuncun-leping/app/front/api/comment/v1"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
)

// Counter is the handler for Comment Controller action.
func (c *ControllerV1) Counter(ctx context.Context, req *v1.CounterReq) (res *v1.CounterRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-comment-Counter")
	defer span.End()

	res = &v1.CounterRes{}
	res.CounterOutput, err = service.Comment().Counter(ctx, req.CounterInput)
	return
}
