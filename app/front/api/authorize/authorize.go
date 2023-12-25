// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package authorize

import (
	"context"

	v1 "github.com/houseme/yuncun-leping/app/front/api/authorize/v1"
)

type IAuthorizeV1 interface {
	Authorize(ctx context.Context, req *v1.AuthorizeReq) (res *v1.AuthorizeRes, err error)
}
