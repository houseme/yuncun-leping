package model

// CommentInput struct
type CommentInput struct {
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
}

// CounterOutput struct
type CounterOutput struct {
	SongsCount      int64 `json:"songs_count" description:"歌曲总数"`
	CommentsCount   int64 `json:"comments_count" description:"评论总数"`
	APIRequestCount int64 `json:"api_request_count" description:"API 请求总数"`
}
