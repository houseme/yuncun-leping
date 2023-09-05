/*
 *  Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  You can obtain one at https://github.com/houseme/yuncun-leping.
 *
 */

// Package domain for domain.
package domain

// SongHotInput struct
type SongHotInput struct {
	Visit string `json:"visit" dc:"访问地址"`
}

// SongHotOutput struct
type SongHotOutput struct {
	List  []*SongItem `json:"list" dc:"歌曲列表"`
	Total int         `json:"total" dc:"歌曲总数"`
}

// SongItem .
type SongItem struct {
	SID   uint64 `json:"sid,string" dc:"歌曲 ID"`
	Title string `json:"title" dc:"歌曲标题"`
}

// SongDetailInput struct
type SongDetailInput struct {
	SID uint64 `json:"sid,string" dc:"歌曲 ID"`
}

// SongDetailOutput struct
type SongDetailOutput struct {
	SID           uint64 `json:"sid,string" dc:"歌曲 ID"`
	Title         string `json:"title" dc:"歌曲标题"`
	Images        string `json:"images" dc:"歌曲图片"`
	Author        string `json:"author" dc:"歌曲作者"`
	Album         string `json:"album" dc:"歌曲专辑"`
	Description   string `json:"description" dc:"歌曲描述"`
	PublishedDate string `json:"publishedDate" dc:"歌曲发布日期"`
	PublishedTime int64  `json:"publishedTime" dc:"歌曲发布时间"`
}

// SongCommentInput struct
type SongCommentInput struct {
	SID uint64 `json:"sid,string" dc:"歌曲 ID"`
}

// SongCommentOutput struct
type SongCommentOutput struct {
	List []*SongCommentItem `json:"list" dc:"歌曲评论列表"`
}

// SongCommentItem struct
type SongCommentItem struct {
	SID           uint64 `json:"sid,string" dc:"歌曲 ID"`
	UserID        uint64 `json:"user_id" dc:"用户 ID"`
	NickName      string `json:"nick_name" dc:"用户昵称"`
	AvatarURL     string `json:"avatarURL" dc:"用户头像"`
	CommentID     uint64 `json:"comment_id" dc:"评论 ID"`
	LikedCount    uint   `json:"liked_count" dc:"评论点赞数"`
	Content       string `json:"content" dc:"评论内容"`
	PublishedDate string `json:"publishedDate" dc:"歌曲发布日期"`
	PublishedTime int64  `json:"publishedTime" dc:"歌曲发布时间"`
}

// QuerySongResponse struct
type QuerySongResponse struct {
	Songs []struct {
		Name        string        `json:"name"`
		ID          int           `json:"id"`
		Position    int           `json:"position"`
		Alias       []interface{} `json:"alias"`
		Status      int           `json:"status"`
		Fee         int           `json:"fee"`
		CopyrightID int           `json:"copyrightId"`
		Disc        string        `json:"disc"`
		No          int           `json:"no"`
		Artists     []struct {
			Name        string        `json:"name"`
			ID          int           `json:"id"`
			PicID       int           `json:"picId"`
			Img1V1Id    int           `json:"img1v1Id"`
			BriefDesc   string        `json:"briefDesc"`
			PicURL      string        `json:"picUrl"`
			Img1V1Url   string        `json:"img1v1Url"`
			AlbumSize   int           `json:"albumSize"`
			Alias       []interface{} `json:"alias"`
			Trans       string        `json:"trans"`
			MusicSize   int           `json:"musicSize"`
			TopicPerson int           `json:"topicPerson"`
		} `json:"artists"`
		Album struct {
			Name        string `json:"name"`
			ID          int    `json:"id"`
			Type        string `json:"type"`
			Size        int    `json:"size"`
			PicID       int64  `json:"picId"`
			BlurPicURL  string `json:"blurPicUrl"`
			CompanyID   int    `json:"companyId"`
			Pic         int64  `json:"pic"`
			PicURL      string `json:"picUrl"`
			PublishTime int64  `json:"publishTime"`
			Description string `json:"description"`
			Tags        string `json:"tags"`
			Company     string `json:"company"`
			BriefDesc   string `json:"briefDesc"`
			Artist      struct {
				Name        string        `json:"name"`
				ID          int           `json:"id"`
				PicID       int           `json:"picId"`
				Img1V1Id    int           `json:"img1v1Id"`
				BriefDesc   string        `json:"briefDesc"`
				PicURL      string        `json:"picUrl"`
				Img1V1Url   string        `json:"img1v1Url"`
				AlbumSize   int           `json:"albumSize"`
				Alias       []interface{} `json:"alias"`
				Trans       string        `json:"trans"`
				MusicSize   int           `json:"musicSize"`
				TopicPerson int           `json:"topicPerson"`
			} `json:"artist"`
			Songs           []interface{} `json:"songs"`
			Alias           []interface{} `json:"alias"`
			Status          int           `json:"status"`
			CopyrightID     int           `json:"copyrightId"`
			CommentThreadID string        `json:"commentThreadId"`
			Artists         []struct {
				Name        string        `json:"name"`
				ID          int           `json:"id"`
				PicID       int           `json:"picId"`
				Img1V1Id    int           `json:"img1v1Id"`
				BriefDesc   string        `json:"briefDesc"`
				PicURL      string        `json:"picUrl"`
				Img1V1Url   string        `json:"img1v1Url"`
				AlbumSize   int           `json:"albumSize"`
				Alias       []interface{} `json:"alias"`
				Trans       string        `json:"trans"`
				MusicSize   int           `json:"musicSize"`
				TopicPerson int           `json:"topicPerson"`
			} `json:"artists"`
			SubType   string      `json:"subType"`
			TransName interface{} `json:"transName"`
			OnSale    bool        `json:"onSale"`
			Mark      int         `json:"mark"`
			Gapless   int         `json:"gapless"`
			DolbyMark int         `json:"dolbyMark"`
			PicIDStr  string      `json:"picId_str"`
		} `json:"album"`
		Starred    bool    `json:"starred"`
		Popularity float64 `json:"popularity"`
		Score      int     `json:"score"`
		StarredNum int     `json:"starredNum"`
		Duration   int     `json:"duration"`
		PlayedNum  int     `json:"playedNum"`
		DayPlays   int     `json:"dayPlays"`
		HearTime   int     `json:"hearTime"`
		SqMusic    struct {
			Name        interface{} `json:"name"`
			ID          int64       `json:"id"`
			Size        int         `json:"size"`
			Extension   string      `json:"extension"`
			Sr          int         `json:"sr"`
			DfsID       int         `json:"dfsId"`
			Bitrate     int         `json:"bitrate"`
			PlayTime    int         `json:"playTime"`
			VolumeDelta float64     `json:"volumeDelta"`
		} `json:"sqMusic"`
		HrMusic              interface{}   `json:"hrMusic"`
		Ringtone             string        `json:"ringtone"`
		Crbt                 interface{}   `json:"crbt"`
		Audition             interface{}   `json:"audition"`
		CopyFrom             string        `json:"copyFrom"`
		CommentThreadID      string        `json:"commentThreadId"`
		RtURL                interface{}   `json:"rtUrl"`
		Ftype                int           `json:"ftype"`
		RtUrls               []interface{} `json:"rtUrls"`
		Copyright            int           `json:"copyright"`
		TransName            interface{}   `json:"transName"`
		Sign                 interface{}   `json:"sign"`
		Mark                 int           `json:"mark"`
		OriginCoverType      int           `json:"originCoverType"`
		OriginSongSimpleData interface{}   `json:"originSongSimpleData"`
		Single               int           `json:"single"`
		NoCopyrightRcmd      interface{}   `json:"noCopyrightRcmd"`
		Mvid                 int           `json:"mvid"`
		BMusic               struct {
			Name        interface{} `json:"name"`
			ID          int64       `json:"id"`
			Size        int         `json:"size"`
			Extension   string      `json:"extension"`
			Sr          int         `json:"sr"`
			DfsID       int         `json:"dfsId"`
			Bitrate     int         `json:"bitrate"`
			PlayTime    int         `json:"playTime"`
			VolumeDelta float64     `json:"volumeDelta"`
		} `json:"bMusic"`
		Rtype  int         `json:"rtype"`
		Rurl   interface{} `json:"rurl"`
		Mp3Url interface{} `json:"mp3Url"`
		HMusic struct {
			Name        interface{} `json:"name"`
			ID          int64       `json:"id"`
			Size        int         `json:"size"`
			Extension   string      `json:"extension"`
			Sr          int         `json:"sr"`
			DfsID       int         `json:"dfsId"`
			Bitrate     int         `json:"bitrate"`
			PlayTime    int         `json:"playTime"`
			VolumeDelta float64     `json:"volumeDelta"`
		} `json:"hMusic"`
		MMusic struct {
			Name        interface{} `json:"name"`
			ID          int64       `json:"id"`
			Size        int         `json:"size"`
			Extension   string      `json:"extension"`
			Sr          int         `json:"sr"`
			DfsID       int         `json:"dfsId"`
			Bitrate     int         `json:"bitrate"`
			PlayTime    int         `json:"playTime"`
			VolumeDelta float64     `json:"volumeDelta"`
		} `json:"mMusic"`
		LMusic struct {
			Name        interface{} `json:"name"`
			ID          int64       `json:"id"`
			Size        int         `json:"size"`
			Extension   string      `json:"extension"`
			Sr          int         `json:"sr"`
			DfsID       int         `json:"dfsId"`
			Bitrate     int         `json:"bitrate"`
			PlayTime    int         `json:"playTime"`
			VolumeDelta float64     `json:"volumeDelta"`
		} `json:"lMusic"`
	} `json:"songs"`
	Equalizers struct {
	} `json:"equalizers"`
	Code int `json:"code"`
}

// CommentResponse struct
type CommentResponse struct {
	IsMusician  bool          `json:"isMusician"`
	UserID      int           `json:"userId"`
	TopComments []interface{} `json:"topComments"`
	MoreHot     bool          `json:"moreHot"`
	HotComments []struct {
		User struct {
			LocationInfo   interface{} `json:"locationInfo"`
			LiveInfo       interface{} `json:"liveInfo"`
			Anonym         int         `json:"anonym"`
			CommonIdentity interface{} `json:"commonIdentity"`
			AvatarDetail   *struct {
				UserType        int    `json:"userType"`
				IdentityLevel   int    `json:"identityLevel"`
				IdentityIconURL string `json:"identityIconUrl"`
			} `json:"avatarDetail"`
			UserType     int         `json:"userType"`
			AvatarURL    string      `json:"avatarUrl"`
			Followed     bool        `json:"followed"`
			Mutual       bool        `json:"mutual"`
			RemarkName   interface{} `json:"remarkName"`
			SocialUserID interface{} `json:"socialUserId"`
			VipRights    *struct {
				Associator *struct {
					VipCode int    `json:"vipCode"`
					Rights  bool   `json:"rights"`
					IconURL string `json:"iconUrl"`
				} `json:"associator"`
				MusicPackage *struct {
					VipCode int    `json:"vipCode"`
					Rights  bool   `json:"rights"`
					IconURL string `json:"iconUrl"`
				} `json:"musicPackage"`
				Redplus           interface{} `json:"redplus"`
				RedVipAnnualCount int         `json:"redVipAnnualCount"`
				RedVipLevel       int         `json:"redVipLevel"`
			} `json:"vipRights"`
			Nickname   string   `json:"nickname"`
			AuthStatus int      `json:"authStatus"`
			ExpertTags []string `json:"expertTags"`
			Experts    *struct {
				Field1 string `json:"1"`
				Field2 string `json:"2"`
			} `json:"experts"`
			VipType int         `json:"vipType"`
			UserID  int64       `json:"userId"`
			Target  interface{} `json:"target"`
		} `json:"user"`
		BeReplied   []interface{} `json:"beReplied"`
		PendantData *struct {
			ID       int    `json:"id"`
			ImageURL string `json:"imageUrl"`
		} `json:"pendantData"`
		ShowFloorComment    interface{} `json:"showFloorComment"`
		Status              int         `json:"status"`
		CommentID           int64       `json:"commentId"`
		Content             string      `json:"content"`
		RichContent         *string     `json:"richContent"`
		ContentResource     interface{} `json:"contentResource"`
		Time                int64       `json:"time"`
		TimeStr             string      `json:"timeStr"`
		NeedDisplayTime     bool        `json:"needDisplayTime"`
		LikedCount          int         `json:"likedCount"`
		ExpressionURL       interface{} `json:"expressionUrl"`
		CommentLocationType int         `json:"commentLocationType"`
		ParentCommentID     int         `json:"parentCommentId"`
		Decoration          struct {
			BubbleID int `json:"bubbleId,omitempty"`
		} `json:"decoration"`
		RepliedMark   interface{} `json:"repliedMark"`
		Grade         interface{} `json:"grade"`
		UserBizLevels interface{} `json:"userBizLevels"`
		IPLocation    struct {
			IP       interface{} `json:"ip"`
			Location string      `json:"location"`
			UserID   *int64      `json:"userId"`
		} `json:"ipLocation"`
		Owner bool `json:"owner"`
		Liked bool `json:"liked"`
	} `json:"hotComments"`
	CommentBanner interface{} `json:"commentBanner"`
	Code          int         `json:"code"`
	Comments      []struct {
		User struct {
			LocationInfo   interface{} `json:"locationInfo"`
			LiveInfo       interface{} `json:"liveInfo"`
			Anonym         int         `json:"anonym"`
			CommonIdentity interface{} `json:"commonIdentity"`
			AvatarDetail   interface{} `json:"avatarDetail"`
			UserType       int         `json:"userType"`
			AvatarURL      string      `json:"avatarUrl"`
			Followed       bool        `json:"followed"`
			Mutual         bool        `json:"mutual"`
			RemarkName     interface{} `json:"remarkName"`
			SocialUserID   interface{} `json:"socialUserId"`
			VipRights      struct {
				Associator *struct {
					VipCode int    `json:"vipCode"`
					Rights  bool   `json:"rights"`
					IconURL string `json:"iconUrl"`
				} `json:"associator"`
				MusicPackage *struct {
					VipCode int    `json:"vipCode"`
					Rights  bool   `json:"rights"`
					IconURL string `json:"iconUrl"`
				} `json:"musicPackage"`
				Redplus           interface{} `json:"redplus"`
				RedVipAnnualCount int         `json:"redVipAnnualCount"`
				RedVipLevel       int         `json:"redVipLevel"`
			} `json:"vipRights"`
			Nickname   string      `json:"nickname"`
			AuthStatus int         `json:"authStatus"`
			ExpertTags interface{} `json:"expertTags"`
			Experts    interface{} `json:"experts"`
			VipType    int         `json:"vipType"`
			UserID     int64       `json:"userId"`
			Target     interface{} `json:"target"`
		} `json:"user"`
		BeReplied []struct {
			User struct {
				LocationInfo   interface{} `json:"locationInfo"`
				LiveInfo       interface{} `json:"liveInfo"`
				Anonym         int         `json:"anonym"`
				CommonIdentity interface{} `json:"commonIdentity"`
				AvatarDetail   interface{} `json:"avatarDetail"`
				UserType       int         `json:"userType"`
				AvatarURL      string      `json:"avatarUrl"`
				Followed       bool        `json:"followed"`
				Mutual         bool        `json:"mutual"`
				RemarkName     interface{} `json:"remarkName"`
				SocialUserID   interface{} `json:"socialUserId"`
				VipRights      interface{} `json:"vipRights"`
				Nickname       string      `json:"nickname"`
				AuthStatus     int         `json:"authStatus"`
				ExpertTags     interface{} `json:"expertTags"`
				Experts        interface{} `json:"experts"`
				VipType        int         `json:"vipType"`
				UserID         int64       `json:"userId"`
				Target         interface{} `json:"target"`
			} `json:"user"`
			BeRepliedCommentID int64       `json:"beRepliedCommentId"`
			Content            string      `json:"content"`
			RichContent        *string     `json:"richContent"`
			Status             int         `json:"status"`
			ExpressionURL      interface{} `json:"expressionUrl"`
			IPLocation         struct {
				IP       interface{} `json:"ip"`
				Location string      `json:"location"`
				UserID   *int64      `json:"userId"`
			} `json:"ipLocation"`
		} `json:"beReplied"`
		PendantData         interface{} `json:"pendantData"`
		ShowFloorComment    interface{} `json:"showFloorComment"`
		Status              int         `json:"status"`
		CommentID           int64       `json:"commentId"`
		Content             string      `json:"content"`
		RichContent         *string     `json:"richContent"`
		ContentResource     interface{} `json:"contentResource"`
		Time                int64       `json:"time"`
		TimeStr             string      `json:"timeStr"`
		NeedDisplayTime     bool        `json:"needDisplayTime"`
		LikedCount          int         `json:"likedCount"`
		ExpressionURL       interface{} `json:"expressionUrl"`
		CommentLocationType int         `json:"commentLocationType"`
		ParentCommentID     int64       `json:"parentCommentId"`
		Decoration          struct {
		} `json:"decoration"`
		RepliedMark   interface{} `json:"repliedMark"`
		Grade         interface{} `json:"grade"`
		UserBizLevels interface{} `json:"userBizLevels"`
		IPLocation    struct {
			IP       interface{} `json:"ip"`
			Location string      `json:"location"`
			UserID   int64       `json:"userId"`
		} `json:"ipLocation"`
		Owner bool `json:"owner"`
		Liked bool `json:"liked"`
	} `json:"comments"`
	Total int  `json:"total"`
	More  bool `json:"more"`
}
