/*
 * Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/yuncun-leping.
 *
 */

package consts

const (
	// TopList 网易云音乐排行榜
	TopList = "https://music.163.com/discover/toplist?id=3778678"

	// SongDetail 网易云音乐歌曲详情
	SongDetail = `https://music.163.com/api/song/detail?ids=[%d]`

	// CommentDetail 网易云音乐歌曲评论
	CommentDetail = `https://music.163.com/api/v1/resource/comments/R_SO_4_%d?limit=20&offset=0`

	// UserAgent .
	UserAgent = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36`

	// HeaderAcceptKey headers accept key
	HeaderAcceptKey = `Accept`

	// HeaderAcceptValue headers accept value
	HeaderAcceptValue = `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8`
)
