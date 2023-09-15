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

package comment

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/houseme/yuncun-leping/app/schedule/internal/consts"
	"github.com/houseme/yuncun-leping/app/schedule/internal/model"
	"github.com/houseme/yuncun-leping/app/schedule/internal/service"
	"github.com/houseme/yuncun-leping/internal/colly"
	"github.com/houseme/yuncun-leping/internal/database/dao"
	"github.com/houseme/yuncun-leping/internal/database/model/do"
	"github.com/houseme/yuncun-leping/internal/database/model/entity"
	"github.com/houseme/yuncun-leping/internal/shared/domain"
	"github.com/houseme/yuncun-leping/utility/helper"
)

var re = regexp.MustCompile(`\s+`) // 匹配一个或多个空格

type sComment struct{}

func init() {
	service.RegisterComment(&sComment{})
}

// QueryOne query one record from table for comment.
func (s *sComment) QueryOne(ctx context.Context, in *model.CommentInput) (out *model.CommentOutput, err error) {
	return
}

// QueryCounter query count from table for comment.
func (s *sComment) QueryCounter(ctx context.Context, in *model.CounterInput) (out *model.CounterOutput, err error) {
	return
}

// initClient init client for comment.
func (s *sComment) initClient(_ context.Context) *gclient.Client {
	return g.Client().SetAgent(consts.UserAgent).SetHeader(consts.HeaderAcceptKey, consts.HeaderAcceptValue).SetTimeout(15 * time.Second)
}

// QuerySongDetail query song detail from table for comment.
func (s *sComment) QuerySongDetail(ctx context.Context, in *domain.SongDetailInput) (out *domain.SongDetailOutput, err error) {
	ctx, span := gtrace.NewSpan(gctx.GetInitCtx(), "tracing-logic-comment-QuerySongDetail")
	defer span.End()

	var response *gclient.Response
	if response, err = s.initClient(ctx).Get(ctx, fmt.Sprintf(consts.SongDetail, in.SID)); err != nil {
		return
	}

	defer func() {
		_ = response.Close()
	}()
	var logger = g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "query song detail http Post request result Response \n", response.Raw())
	var resp *domain.QuerySongResponse
	if err = gjson.New(response.ReadAll()).Scan(&resp); err != nil {
		return
	}

	if resp == nil || resp.Code != 200 {
		logger.Debug(ctx, "query song detail response is empty")
		return
	}

	if resp.Songs == nil || len(resp.Songs) < 1 {
		logger.Debug(ctx, "query song detail songs info is empty")
		return
	}

	out = &domain.SongDetailOutput{
		SID:           in.SID,
		Title:         resp.Songs[0].Name,
		Images:        strings.ReplaceAll(resp.Songs[0].Album.PicURL, "http://", "https://"),
		Author:        resp.Songs[0].Artists[0].Name,
		Album:         resp.Songs[0].Album.Name,
		Description:   "歌手：" + resp.Songs[0].Artists[0].Name + " 专辑：" + resp.Songs[0].Album.Name,
		PublishedTime: resp.Songs[0].Album.PublishTime,
		PublishedDate: gtime.NewFromTimeStamp(resp.Songs[0].Album.PublishTime).Format("Y-m-d H:i:s"),
	}

	if gstr.Trim(out.Title) == "" || gstr.Trim(out.Author) == "" || gstr.Trim(out.Album) == "" {
		logger.Debugf(ctx, "query song detail title author album description is empty content: %+v", out)
		return
	}
	var lastID int64
	if lastID, err = dao.Songs.Ctx(ctx).OmitEmpty().Unscoped().InsertAndGetId(do.Songs{
		SongId:        resp.Songs[0].ID,
		Title:         resp.Songs[0].Name,
		Images:        strings.ReplaceAll(resp.Songs[0].Album.PicURL, "http://", "https://"),
		Author:        resp.Songs[0].Artists[0].Name,
		Album:         resp.Songs[0].Album.Name,
		Description:   "歌手：" + resp.Songs[0].Artists[0].Name + " 专辑：" + resp.Songs[0].Album.Name,
		PublishTime:   resp.Songs[0].Album.PublishTime,
		PublishedDate: gtime.NewFromTimeStamp(resp.Songs[0].Album.PublishTime),
	}); err != nil {
		return
	}
	logger.Debugf(ctx, "query song detail insert song id: %d", lastID)

	return
}

// QuerySongComment query song comment from table for comment.
func (s *sComment) QuerySongComment(ctx context.Context, in *domain.SongCommentInput) (out *domain.SongCommentOutput, err error) {
	ctx, span := gtrace.NewSpan(gctx.GetInitCtx(), "tracing-logic-comment-QuerySongComment")
	defer span.End()

	var response *gclient.Response
	if response, err = s.initClient(ctx).Get(ctx, fmt.Sprintf(consts.CommentDetail, in.SID)); err != nil {
		return
	}

	defer func() {
		_ = response.Close()
	}()
	var logger = g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "query song comment http Post request result Response \n", response.Raw())
	var resp *domain.CommentResponse
	if err = gjson.New(response.ReadAll()).Scan(&resp); err != nil {
		return
	}

	if resp == nil || resp.Code != 200 {
		logger.Debug(ctx, "query song comment response is empty")
		return
	}

	if resp.HotComments == nil || len(resp.HotComments) < 1 {
		logger.Debug(ctx, "query song comment hot comments is empty")
		return
	}
	out = &domain.SongCommentOutput{
		List: []*domain.SongCommentItem{},
	}
	for i := 0; i < len(resp.HotComments); i++ {
		var (
			commentEntity = (*entity.HotComments)(nil)
			comment       = resp.HotComments[i]
		)
		if err = dao.HotComments.Ctx(ctx).Scan(&commentEntity, do.HotComments{CommentId: comment.CommentID}); err != nil {
			logger.Errorf(ctx, "query song comment failed error: %+v", err)
			continue
		}
		var (
			likedCount = uint(comment.LikedCount)
			lastID     int64
		)
		out.List = append(out.List, &domain.SongCommentItem{
			SID:           in.SID,
			CommentID:     uint64(comment.CommentID),
			Content:       re.ReplaceAllString(comment.Content, " "),
			LikedCount:    likedCount,
			UserID:        uint64(comment.User.UserID),
			NickName:      comment.User.Nickname,
			AvatarURL:     comment.User.AvatarURL,
			PublishedTime: comment.Time,
			PublishedDate: gtime.NewFromTimeStamp(comment.Time).Format("Y-m-d H:i:s"),
		})

		if commentEntity != nil {
			if likedCount != commentEntity.LikedCount {
				if lastID, err = dao.HotComments.Ctx(ctx).OmitEmpty().Unscoped().UpdateAndGetAffected(do.HotComments{
					LikedCount:    likedCount,
					PublishTime:   comment.Time,
					UserId:        comment.User.UserID,
					Nickname:      comment.User.Nickname,
					PublishedDate: gtime.NewFromTimeStamp(comment.Time),
					ModifyTime:    gtime.Now(),
				}, do.HotComments{Id: commentEntity.Id}); err != nil {
					logger.Errorf(ctx, "query song comment update failed error: %+v", err)
					continue
				}
				logger.Debugf(ctx, "query song comment entity update lastID: %d", commentEntity.Id)
			}
			continue
		}

		if lastID, err = dao.HotComments.Ctx(ctx).OmitEmpty().Unscoped().InsertAndGetId(do.HotComments{
			SongId:        in.SID,
			UserId:        comment.User.UserID,
			Nickname:      comment.User.Nickname,
			AvatarUrl:     comment.User.AvatarURL,
			CommentId:     comment.CommentID,
			LikedCount:    likedCount,
			ModifyTime:    gtime.Now(),
			PublishTime:   comment.Time,
			PublishedDate: gtime.NewFromTimeStamp(comment.Time),
			Content:       re.ReplaceAllString(comment.Content, " "),
		}); err != nil {
			logger.Errorf(ctx, "query song comment insert failed error: %+v", err)
			continue
		}
		logger.Debugf(ctx, "query song comment entity insert lastID: %d", lastID)
	}
	logger.Debugf(ctx, "query song comment entity insert end total: %d", resp.Total)
	return
}

// TopList query top list from table for comment.
func (s *sComment) TopList(ctx context.Context) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-comment-TopList")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Info(ctx, "cron job logic top list running start")
	list, err := colly.GetList(ctx, consts.TopList)
	if err != nil {
		logger.Errorf(ctx, "cron job top list running colly get list failed error: %+v", err)
		return err
	}
	if list == nil || len(list) < 1 {
		logger.Debug(ctx, "cron job top list colly list is empty")
		return
	}
	var wg sync.WaitGroup
	for i, item := range list {
		logger.Infof(ctx, "cron job top list colly list ID: %d item: %+v", i, item)
		go s.QuerySong(ctx, item.SID, &wg)
		go s.QueryComment(ctx, item.SID, &wg)
	}
	wg.Wait()
	logger.Info(ctx, "cron job top list running end")
	return
}

// QuerySong query song from table for comment.
func (s *sComment) QuerySong(ctx context.Context, sid uint64, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	var (
		traceID = gtrace.GetTraceID(ctx)
		span    *gtrace.Span
		logger  = g.Log(helper.Helper().Logger(ctx))
	)
	ctx, span = gtrace.NewSpan(gctx.GetInitCtx(), "tracing-logic-comment-QuerySong")
	span.End()

	var err error
	if ctx, err = gtrace.WithTraceID(ctx, traceID); err != nil {
		logger.Errorf(ctx, "cron job async query song failed error: %+v", err)
		return
	}

	var songEntity = (*entity.Songs)(nil)
	if err = dao.Songs.Ctx(ctx).Scan(&songEntity, do.Songs{SongId: sid}); err != nil {
		logger.Errorf(ctx, "cron job async query song failed error: %+v", err)
		return
	}

	if songEntity == nil {
		var songDetail *domain.SongDetailOutput
		if songDetail, err = s.QuerySongDetail(ctx, &domain.SongDetailInput{SID: sid}); err != nil {
			logger.Errorf(ctx, "cron job async query song detail failed error: %+v", err)
			return
		}
		logger.Debug(ctx, "cron job async song detail:", songDetail)
	}
	logger.Debugf(ctx, "cron job async song end sid: %d", sid)
	return
}

// QueryComment query comment from table for comment.
func (s *sComment) QueryComment(ctx context.Context, sid uint64, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	var (
		traceID = gtrace.GetTraceID(ctx)
		span    *gtrace.Span
		logger  = g.Log(helper.Helper().Logger(ctx))
	)
	ctx, span = gtrace.NewSpan(gctx.GetInitCtx(), "tracing-logic-comment-QueryComment")
	span.End()
	var err error
	if ctx, err = gtrace.WithTraceID(ctx, traceID); err != nil {
		logger.Errorf(ctx, "cron job async query comment failed error: %+v", err)
		return
	}

	var songComment *domain.SongCommentOutput
	if songComment, err = s.QuerySongComment(ctx, &domain.SongCommentInput{SID: sid}); err != nil {
		logger.Errorf(ctx, "cron job async query song comment failed error: %+v", err)
		return
	}
	logger.Debugf(ctx, "cron job async song comment end sid: %d hot comment length: %d", sid, len(songComment.List))
	return
}
