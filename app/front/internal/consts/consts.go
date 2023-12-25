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
	// DefaultLogger .
	DefaultLogger = ""

	// AuthorizationHeaderKey 授权头
	AuthorizationHeaderKey = "authorization"
	// AuthorizationTypeBearer 授权类型
	AuthorizationTypeBearer = "Bearer"

	// AuthTypeAPIKey 授权类型 ApiKey 授权类型 账户密码授权类型
	AuthTypeAPIKey = "api_key"
	// AuthTypePassword 授权类型 账户密码授权类型
	AuthTypePassword = "password"

	// APIKeyExpireTime ApiKey 授权有效期时间 单位秒
	APIKeyExpireTime = 7200

	// PasswordExpireTime 账号密码授权有效期时间 单位秒
	PasswordExpireTime = 7200

	// AccessTokenExpireTime 访问令牌有效期时间 单位秒
	AccessTokenExpireTime = 7200

	// RefreshTokenExpireTime 刷新令牌有效期时间 单位秒
	RefreshTokenExpireTime = 864000

	// TokenExpireTime 有效期时间 单位秒
	TokenExpireTime = 864000

	// AuthTypeDefault 授权类型 0 默认，认证类型 100 app key，200 app id
	AuthTypeDefault = 0
	// AuthTypeAppKey 授权类型 100 app key，200 app id
	AuthTypeAppKey = 100
	// AuthTypeAppID 授权类型 100 app key，200 app id
	AuthTypeAppID = 200
)

const (
	// RedirectType 0 默认 100 歌曲，200 歌词
	RedirectType = 0
	// RedirectTypeSong 歌曲
	RedirectTypeSong = 100
	// RedirectTypeLyric 歌词
	RedirectTypeLyric = 200

	// MusicContentType mp3 歌曲 lyric 歌词
	MusicContentType = "mp3"
	// LyricContentType 歌词
	LyricContentType = "lyric"
)
