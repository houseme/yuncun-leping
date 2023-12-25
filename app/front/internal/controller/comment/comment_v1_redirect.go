package comment

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/yuncun-leping/app/front/api/comment/v1"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
)

// Redirect is the handler for Comment Controller action.
func (c *ControllerV1) Redirect(ctx context.Context, req *v1.RedirectReq) (res *v1.RedirectRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-comment-Redirect")
	defer span.End()

	res = &v1.RedirectRes{}
	if res.RedirectOutput, err = service.Comment().Redirect(ctx, req.RedirectInput); err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(res.RedirectURL)
	return
}
