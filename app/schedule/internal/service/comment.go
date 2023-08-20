// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/houseme/yuncun-leping/app/schedule/internal/model"
)

type (
	IComment interface {
		// QueryOne query one record from table for comment.
		QueryOne(ctx context.Context, in *model.CommentInput) (out *model.CommentOutput, err error)
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
