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

package env

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// SnowflakeEnv .
type SnowflakeEnv struct {
	datacenter int64
	worker     int64
	config     map[string]*gvar.Var
}

// Datacenter .
func (s *SnowflakeEnv) Datacenter(_ context.Context) int64 {
	return s.datacenter
}

// Worker .
func (s *SnowflakeEnv) Worker(_ context.Context) int64 {
	return s.worker
}

// String .
func (s *SnowflakeEnv) String(_ context.Context) string {
	return `{"datacenter":"` + gconv.String(s.datacenter) + `","worker":"` + gconv.String(s.worker) + `"}`
}

// NewSnowflakeEnv .
func NewSnowflakeEnv(ctx context.Context) (*SnowflakeEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewSnowflakeEnv")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, "snowflake")
	if err != nil {
		return nil, gerror.Wrap(err, "config snowflake get failed")
	}
	if v.IsNil() || v.IsEmpty() {
		return nil, gerror.New("config snowflake is empty")
	}

	var (
		config = v.MapStrVar()
		env    = &SnowflakeEnv{
			worker:     config["worker"].Int64(),
			datacenter: config["datacenter"].Int64(),
			config:     config,
		}
	)
	return env, nil
}
