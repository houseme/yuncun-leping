package main

import (
	_ "github.com/houseme/yuncun-leping/app/schedule/internal/packed"

	_ "github.com/houseme/yuncun-leping/app/schedule/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/houseme/yuncun-leping/app/schedule/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
