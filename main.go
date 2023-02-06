package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"goFrameMall/internal/cmd"
	_ "goFrameMall/internal/logic"
	_ "goFrameMall/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
