// Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// You can obtain one at https://github.com/houseme/yuncun-leping.

// Package comment for comment.
package comment

import (
	"context"
	"fmt"

	"github.com/houseme/yuncun-leping/app/front/internal/model"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
	"github.com/houseme/yuncun-leping/internal/database/dao"
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
	out = &model.CounterOutput{
		SongsCount:      0,
		CommentsCount:   0,
		APIRequestCount: 0,
	}
	if out.SongsCount, err = dao.Songs.Ctx(ctx).Count(); err != nil {
		return
	}
	if out.CommentsCount, err = dao.HotComments.Ctx(ctx).Count(); err != nil {
		return
	}
	return
}

// Redirect to music.
func (s *sComment) Redirect(ctx context.Context, in *model.RedirectInput) (out *model.RedirectOutput, err error) {
	out = &model.RedirectOutput{
		RedirectURL: fmt.Sprintf("https://music.163.com/song/media/outer/url?id=%d.mp3", in.SongID),
	}
	return
}
