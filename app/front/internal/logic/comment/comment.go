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

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/houseme/yuncun-leping/app/front/internal/consts"
	"github.com/houseme/yuncun-leping/app/front/internal/model"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
	"github.com/houseme/yuncun-leping/internal/database/dao"
	"github.com/houseme/yuncun-leping/internal/database/model/do"
	"github.com/houseme/yuncun-leping/utility/cache"
	"github.com/houseme/yuncun-leping/utility/env"
	"github.com/houseme/yuncun-leping/utility/helper"
)

type sComment struct{}

func init() {
	service.RegisterComment(&sComment{})
}

// Home query one record from table for comment.
func (s *sComment) Home(ctx context.Context, in *model.HomeInput) (out *model.HomeOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-logic-Home")
	defer span.End()

	var (
		query  = `SELECT t3.song_id,t3.title, t3.images, t3.author, t3.album, t3.description, '' as 'mp3_url',t3.published_date as publish_date,t1.comment_id,t1.user_id AS comment_user_id,t1.nickname AS comment_nickname,t1.avatar_url AS comment_avatar,t1.liked_count AS comment_liked_count,t1.content AS comment_content,t1.published_date AS comment_published_date FROM hot_comments t1 JOIN (SELECT ROUND(RAND() * ((SELECT MAX(id) FROM hot_comments) - ( SELECT MIN(id) FROM hot_comments )) + (SELECT MIN(id) FROM hot_comments )) AS id ) t2 JOIN songs t3 ON t1.song_id = t3.song_id WHERE t1.id = t2.id LIMIT 1;`
		logger = g.Log(helper.Helper().Logger(ctx))
	)
	logger.Debug(ctx, "comment home logic start ")
	if err = g.DB().GetScan(ctx, &out, query); err != nil {
		return
	}

	var (
		now          = gtime.Now()
		lastID       int64
		counterValue int64
	)
	if lastID, err = g.Redis(cache.CounterConn(ctx)).Incr(ctx, cache.DefaultCounterKey(ctx)); err != nil {
		return
	}
	logger.Debugf(ctx, "home Redis incr request log last id: %d", lastID)
	counterValue = lastID
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var requestID = helper.Helper().InitTrxID(ctx, in.AuthAppNo)
		if lastID, err = dao.RequestLog.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(do.RequestLog{
			RequestId:   requestID,
			AppNo:       in.AuthAppNo,
			YearTime:    now.Year(),
			MonthTime:   now.Month(),
			DayTime:     now.Day(),
			UserAgent:   in.UserAgent,
			Referer:     in.Referer,
			Origin:      in.Origin,
			Path:        in.Path,
			RequestUri:  in.RequestURI,
			RequestIp:   in.ClientIP,
			RequestTime: gtime.NewFromTimeStamp(g.RequestFromCtx(ctx).EnterTime),
		}); err != nil {
			return err
		}
		logger.Debugf(ctx, "home insert request log last id: %d", lastID)
		var data = do.ResponseLog{
			RequestId:    requestID,
			AppNo:        in.AuthAppNo,
			YearTime:     now.Year(),
			MonthTime:    now.Month(),
			DayTime:      now.Day(),
			RequestIp:    in.ClientIP,
			RequestTime:  gtime.NewFromTimeStamp(g.RequestFromCtx(ctx).EnterTime),
			ResponseTime: gtime.Now(),
			CounterValue: counterValue,
		}

		if out != nil {
			data.SongId = out.SongID
			data.CommentId = out.CommentID
			data.CommentUserId = out.CommentUserID
			data.CommentLikedCount = out.CommentLikedCount
			data.Title = out.Title
			data.CommentContent = out.CommentContent
		}
		if lastID, err = dao.ResponseLog.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(data); err != nil {
			return err
		}
		logger.Debugf(ctx, "home insert response log last id: %d,requestID: %d", lastID, requestID)
		return nil
	}); err != nil {
		return
	}

	if out != nil {
		var appEnv *env.AppEnv
		if appEnv, err = env.New(ctx); err != nil {
			return
		}
		out.Mp3URL = appEnv.Site() + "api.v1/front/music/" + gconv.String(out.SongID) + "/" + consts.MusicContentType
		// "https://music.163.com/api/song/media?id=" + gconv.String(out.SongID)
		out.LyricURL = appEnv.Site() + "api.v1/front/music/" + gconv.String(out.SongID) + "/" + consts.LyricContentType
	}

	if in.AuthAppNo > 0 {
		if lastID, err = g.Redis(cache.DefaultConn(ctx)).Incr(ctx, cache.DefaultCounterByAppKey(ctx, in.AuthAppNo)); err != nil {
			return
		}
		logger.Debugf(ctx, "home Redis incr request log last id: %d app no: %d", lastID, in.AuthAppNo)
	}
	logger.Debug(ctx, "comment home logic success end")
	return
}

// Counter query count from table for comment.
func (s *sComment) Counter(ctx context.Context, in *model.CounterInput) (out *model.CounterOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-logic-Counter")
	defer span.End()

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

	var (
		value    *gvar.Var
		redisKey = cache.DefaultCounterKey(ctx)
	)
	if in.AuthAppNo > 0 {
		redisKey = cache.DefaultCounterByAppKey(ctx, in.AuthAppNo)
	}
	if value, err = g.Redis(cache.CounterConn(ctx)).Get(ctx, redisKey); err != nil {
		return
	}
	if value != nil && !value.IsNil() && !value.IsEmpty() {
		out.APIRequestCount = value.Int()
		return
	}

	return
}

// Redirect to music.
func (s *sComment) Redirect(ctx context.Context, in *model.RedirectInput) (out *model.RedirectOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-logic-Redirect")
	defer span.End()

	out = &model.RedirectOutput{
		RedirectURL: fmt.Sprintf("https://music.163.com/song/media/outer/url?id=%d.mp3", in.SongID),
	}
	if in.ContentType == "lyric" {
		out.RedirectURL = fmt.Sprintf("https://music.163.com/api/song/media?id=%d", in.SongID)
	}
	return
}
