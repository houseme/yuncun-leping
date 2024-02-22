package comment

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/yuncun-leping/app/front/api/comment/v1"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
	"github.com/houseme/yuncun-leping/utility/helper"
)

// CompatibleCount compatible count
func (c *ControllerV1) CompatibleCount(ctx context.Context, req *v1.CompatibleCountReq) (res *v1.CompatibleCountRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-comment-Counter")
	defer span.End()

	var (
		r      = g.RequestFromCtx(ctx)
		logger = g.Log(helper.Helper().Logger(ctx))
	)

	res = &v1.CompatibleCountRes{}
	if res.CounterOutput, err = service.Comment().Counter(ctx, req.CounterInput); err != nil {
		logger.Errorf(ctx, "service.Comment().Counter failed error: %v", err)
		r.Response.WriteStatusExit(503, "系统繁忙，请稍后重试")
	}
	logger.Debug(r.GetCtx(), "comment home logic success out:", res.CounterOutput)
	r.Response.WriteJsonExit(res.CounterOutput)
	return
}
