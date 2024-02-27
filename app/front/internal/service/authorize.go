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
	IAuthorize interface {
		// Authorization app authorization
		Authorization(ctx context.Context, in *model.AuthorizeInput) (out *model.AuthorizeOutput, err error)
	}
)

var (
	localAuthorize IAuthorize
)

func Authorize() IAuthorize {
	if localAuthorize == nil {
		panic("implement not found for interface IAuthorize, forgot register?")
	}
	return localAuthorize
}

func RegisterAuthorize(i IAuthorize) {
	localAuthorize = i
}
