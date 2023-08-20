package comment

import (
	"context"

	"github.com/houseme/yuncun-leping/app/schedule/internal/model"
	"github.com/houseme/yuncun-leping/app/schedule/internal/service"
)

type sComment struct{}

func init() {
	service.RegisterComment(&sComment{})
}

// QueryOne query one record from table for comment.
func (s *sComment) QueryOne(ctx context.Context, in *model.CommentInput) (out *model.CommentOutput, err error) {
	return
}
