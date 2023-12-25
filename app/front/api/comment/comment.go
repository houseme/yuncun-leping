// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package comment

import (
	"context"

	v1 "github.com/houseme/yuncun-leping/app/front/api/comment/v1"
)

type ICommentV1 interface {
	Home(ctx context.Context, req *v1.HomeReq) (res *v1.HomeRes, err error)
	Redirect(ctx context.Context, req *v1.RedirectReq) (res *v1.RedirectRes, err error)
	Counter(ctx context.Context, req *v1.CounterReq) (res *v1.CounterRes, err error)
	Compatible(ctx context.Context, req *v1.CompatibleReq) (res *v1.CompatibleRes, err error)
}
