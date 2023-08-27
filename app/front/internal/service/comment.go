// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/houseme/yuncun-leping/app/front/internal/model"
)

type (
	IComment interface {
		// One query one record from table for comment.
		Home(ctx context.Context, in *model.CommentInput) (out *model.CommentOutput, err error)
		// Counter query count from table for comment.
		Counter(ctx context.Context, in *model.CounterInput) (out *model.CounterOutput, err error)
		// Redirect to music.
		Redirect(ctx context.Context, in *model.RedirectInput) (out *model.RedirectOutput, err error)
	}
)

var (
	localComment IComment
)

func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

func RegisterComment(i IComment) {
	localComment = i
}
