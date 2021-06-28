package app

import (
	"github.com/itsfinn/autossh/utils"
)

func showVersion() {
	utils.Logln("autossh " + Version + " Build " + Build + "。")
	utils.Logln("由 Lenbo 编写，项目地址：https://github.com/islenbo/autossh。")
}
