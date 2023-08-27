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
