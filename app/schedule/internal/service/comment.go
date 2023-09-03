// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/houseme/yuncun-leping/app/schedule/internal/model"
	"github.com/houseme/yuncun-leping/internal/shared/domain"
)

type (
	IComment interface {
		// QueryOne query one record from table for comment.
		QueryOne(ctx context.Context, in *model.CommentInput) (out *model.CommentOutput, err error)
		// QueryCounter query count from table for comment.
		QueryCounter(ctx context.Context, in *model.CounterInput) (out *model.CounterOutput, err error)
		// QuerySongDetail query song detail from table for comment.
		QuerySongDetail(ctx context.Context, in *domain.SongDetailInput) (out *domain.SongDetailOutput, err error)
		// QuerySongComment query song comment from table for comment.
		QuerySongComment(ctx context.Context, in *domain.SongCommentInput) (out *domain.SongCommentOutput, err error)
		// TopList query top list from table for comment.
		TopList(ctx context.Context) (err error)
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
