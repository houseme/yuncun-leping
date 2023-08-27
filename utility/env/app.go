package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

// New  创建 APP 环境
func New(ctx context.Context) (*AppEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-New")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, "app")
	if err != nil {
		return nil, gerror.Wrap(err, "config app get failed")
	}
	if v.IsNil() || v.IsEmpty() {
		return nil, gerror.New("config app is empty")
	}
	var config = v.MapStrStr()
	return &AppEnv{
		config:         config,
		env:            config["env"],
		environment:    config["environment"],
		version:        config["version"],
		jaegerEndpoint: config["jaegerEndpoint"],
		endpoint:       config["endpoint"],
		traceToken:     config["traceToken"],
		uploadPath:     config["uploadPath"],
		visitPath:      config["visitPath"],
		site:           config["site"],
		roleModel:      config["roleModel"],
		frontSite:      config["frontSite"],
	}, nil
}

// AppEnv .
type AppEnv struct {
	config         map[string]string
	env            string
	environment    string
	version        string
	jaegerEndpoint string
	endpoint       string
	traceToken     string
	uploadPath     string
	visitPath      string
	site           string
	roleModel      string
	frontSite      string
}

// Env .
func (a *AppEnv) Env() string {
	return a.env
}

// Environment .
func (a *AppEnv) Environment() string {
	return a.environment
}

// Version .
func (a *AppEnv) Version() string {
	return a.version
}

// JaegerEndpoint .
func (a *AppEnv) JaegerEndpoint() string {
	return a.jaegerEndpoint
}

// Endpoint .
func (a *AppEnv) Endpoint() string {
	return a.endpoint
}

// TraceToken .
func (a *AppEnv) TraceToken() string {
	return a.traceToken
}

// Config .获取配置信息
func (a *AppEnv) Config() map[string]string {
	return a.config
}

// UploadPath .	上传路径
func (a *AppEnv) UploadPath() string {
	return a.uploadPath
}

// VisitPath file server 访问路径
func (a *AppEnv) VisitPath() string {
	return a.visitPath
}

// Site .网站名称
func (a *AppEnv) Site() string {
	return a.site
}

// RoleModel .
func (a *AppEnv) RoleModel() string {
	return a.roleModel
}

// FrontSite .
func (a *AppEnv) FrontSite() string {
	return a.frontSite
}

// String
func (a *AppEnv) String() string {
	return `{"env":"` + a.env + `","environment":"` + a.environment + `","version":"` + a.version +
		`","jaegerEndpoint":"` + a.jaegerEndpoint + `","endpoint":"` + a.endpoint +
		`","uploadPath":"` + a.uploadPath + `","visitPath":"` + a.visitPath +
		`","site":"` + a.site + `","roleModel":"` + a.roleModel + `","frontSite":"` + a.frontSite + `","traceToken":"` + a.traceToken + `"}`
}
