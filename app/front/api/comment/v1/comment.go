package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/yuncun-leping/app/front/internal/model"
)

// Req struct
type Req struct {
	g.Meta `path:"/" tags:"Comment" method:"get" summary:"music comment"`
	*model.CommentInput
}

// Res struct
type Res struct {
	*model.CommentOutput
}
