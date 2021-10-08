package common

import (
	"github.com/youthlin/t"
	"github.com/youthlin/z"
)

func MustInit() {
	initConfig()
	initLog()
	initI18n()
}

func initLog() {
	z.SetConfig(&c.Logs)
}

func initI18n() {
	t.Load(c.LangPath)
}
