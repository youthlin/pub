package common

import (
	"github.com/youthlin/z"
)

func MustInit() {
	initConfig()
	initLog()
}

func initLog() {
	z.SetConfig(&c.Logs)
}
