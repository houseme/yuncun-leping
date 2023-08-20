package main

import (
	_ "github.com/houseme/yuncun-leping/app/front/internal/packed"

	_ "github.com/houseme/yuncun-leping/app/front/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/houseme/yuncun-leping/app/front/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
