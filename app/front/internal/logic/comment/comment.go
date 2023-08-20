package comment

import (
	"context"

	"github.com/houseme/yuncun-leping/app/front/internal/model"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
)

type sComment struct{}

func init() {
	service.RegisterComment(&sComment{})
}

// One query one record from table for comment.
func (s *sComment) One(ctx context.Context, in *model.CommentInput) (out *model.CommentOutput, err error) {
	return
}

// Counter query count from table for comment.
func (s *sComment) Counter(ctx context.Context, in *model.CounterInput) (out *model.CounterOutput, err error) {
	return
}
