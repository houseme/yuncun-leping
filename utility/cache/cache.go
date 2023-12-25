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

package cache

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

const (
	defaultConn = "default"

	// defaultExpiration is the default expiration time of cache item.
	defaultExpiration = 120

	// noExpiration means the cache items have no expiration time.
	noExpiration = 0

	// counter key
	counterKey = "yc:counter"

	// token 链接
	centerAccessTokenConn = "center_access_token"

	// 登陆授权生成的 token
	centerAccessTokenKey = "yc:center_access_token_key_"

	// centerAuthorizationKey 密钥生成的 token
	centerAuthorizationKey = "yc:center_authorization_key_"
)

// DefaultConn returns the default connection of cache.
func DefaultConn(_ context.Context) string {
	return defaultConn
}

// DefaultExpiration returns the default expiration of cache item.
func DefaultExpiration(_ context.Context) int {
	return defaultExpiration
}

// NoExpiration returns the no expiration of cache item.
func NoExpiration(_ context.Context) int {
	return noExpiration
}

// CounterKey returns the counter key of cache item.
func CounterKey(_ context.Context) string {
	return counterKey
}

// CenterAccessTokenConn returns the Redis connection
func CenterAccessTokenConn(_ context.Context) string {
	return centerAccessTokenConn
}

// CenterAccessTokenKey returns the Redis access token key
func CenterAccessTokenKey(_ context.Context, accessToken string) string {
	return centerAccessTokenKey + accessToken
}

// CenterAuthorizationKey returns the Redis authorization key
func CenterAuthorizationKey(_ context.Context, authorization string) string {
	return centerAuthorizationKey + authorization
}

// CounterByAppKey returns the counter key of cache item.
func CounterByAppKey(_ context.Context, appKey uint64) string {
	return counterKey + ":" + gconv.String(appKey)
}
