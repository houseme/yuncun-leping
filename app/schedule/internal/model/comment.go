package model

// CommentInput struct
type CommentInput struct {
}

// CommentOutput struct
type CommentOutput struct {
	SongID int64 `json:"song_id" description:"歌曲 ID"`
}
