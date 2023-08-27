/*
 *   Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 *
 *   You can obtain one at https://github.com/houseme/yuncun-leping.
 *
 */

package model

// CommentInput struct
type CommentInput struct {
	*AuthBase `json:"-"`
}

// CommentOutput struct
type CommentOutput struct {
	SongID               int64  `json:"song_id" description:"歌曲 ID"`
	Title                string `json:"title" description:"歌曲标题"`
	Images               string `json:"images" description:"歌曲封面"`
	Author               string `json:"author" description:"歌曲作者"`
	Summary              string `json:"summary" description:"歌曲简介"`
	Album                string `json:"album" description:"歌曲专辑"`
	Mp3URL               string `json:"mp3_url" description:"歌曲地址"`
	LyricURL             string `json:"lyric_url" description:"歌词地址"`
	PublishDate          string `json:"publish_date" description:"歌曲发布日期"`
	CommentID            int64  `json:"comment_id" description:"评论 ID"`
	CommentUserID        int64  `json:"comment_user_id" description:"评论用户 ID"`
	CommentNickName      string `json:"comment_nick_name" description:"评论用户昵称"`
	CommentAvatar        string `json:"comment_avatar" description:"评论用户头像"`
	CommentContent       string `json:"comment_content" description:"评论内容"`
	CommentLikedCount    int64  `json:"comment_liked_count" description:"评论点赞数"`
	CommentPublishedDate string `json:"comment_published_date" description:"评论发布日期"`
}

// CounterInput struct
type CounterInput struct {
	*AuthBase `json:"-"`
}

// CounterOutput struct
type CounterOutput struct {
	SongsCount      int `json:"songs_count" description:"歌曲总数"`
	CommentsCount   int `json:"comments_count" description:"评论总数"`
	APIRequestCount int `json:"api_request_count" description:"API 请求总数"`
}

// RedirectInput struct
type RedirectInput struct {
	*AuthBase   `json:"-"`
	SongID      uint64 `json:"song_id" description:"歌曲 ID"`
	ContentType string `json:"content_type" description:"内容类型 mp3|lyric" v:"required|in:mp3,lyric#内容类型不能为空 | 内容类型必须在 mp3,lyric 中"`
}

// RedirectOutput struct
type RedirectOutput struct {
	RedirectURL string `json:"redirect_url" description:"重定向地址"`
}
