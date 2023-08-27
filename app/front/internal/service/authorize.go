// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IAuthorize interface{}
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
