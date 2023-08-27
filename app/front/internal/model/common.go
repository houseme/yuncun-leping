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

// DefaultHandlerResponse .
type DefaultHandlerResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Time    int64       `json:"time,string"` // 返回当前响应时间
	TraceID string      `json:"traceID"`     // 请求唯一标识
}

// AuthorizationToken is the golang structure for AuthorizationToken.
type AuthorizationToken struct {
	AuthToken string `json:"authToken" dc:"认证 token"`
	AuthTime  int64  `json:"authTime" dc:"认证时间"`
	AuthAppNo uint64 `json:"authAppNo" dc:"认证应用编号"`
	AuthType  uint   `json:"authType" dc:"认证类型 100 app key，200 app id"`
}

// AuthBase is the base struct for auth.
type AuthBase struct {
	AuthAppNo uint64 `json:"authAppNo" dc:"认证应用编号"`
	AuthType  uint   `json:"authType" dc:"认证类型 100 app key，200 app id"`
}
