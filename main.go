package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/logs"
	"github.com/youthlin/pub/common"
	"github.com/youthlin/z"
)

var log = logs.GetLogger()

func main() {
	common.MustInit()
	server := startWeb()
	graceffulyShutdown(server)
}

func startWeb() *http.Server {
	cfg := common.Config()
	if cfg == nil {
		panic(errors.Errorf("app config is nil"))
	}
	server := &http.Server{
		Addr:    cfg.Web.Addr,
		Handler: newEngine(&cfg.Web),
	}
	go server.ListenAndServe()
	return server
}

func newEngine(cfg *common.WebConfig) *gin.Engine {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// Access
	var h []gin.HandlerFunc
	for _, a := range cfg.AccessLog {
		var w io.Writer = a
		if a.Type == z.Console { // return Stdout/Stderr so that can show colorful http status/method
			if a.File.Filename == z.Stderr {
				w = os.Stderr
			} else {
				w = os.Stdout
			}
		}
		h = append(h, gin.LoggerWithWriter(w))
	}
	r.Use(h...)
	// Recovery
	r.Use(gin.RecoveryWithWriter(cfg.ErrorLog.MultiWriter()))
	// routers
	register(r)
	return r
}

// graceffulyShutdown 优雅停机
// https://mp.weixin.qq.com/s/mnR_wADjyHtyjdkSwcnkSw
func graceffulyShutdown(server *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	<-ctx.Done() // 监听中断信号
	stop()       // 重置 os.Interrupt 的默认行为, 再次按下 ^C 会直接退出
	log.Info("shutting down graceffuly, press Ctrl+C again to force exist immediately")

	// 最多等待 5 分钟
	timeOutCtx, cancal := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancal()
	if err := server.Shutdown(timeOutCtx); err != nil {
		log.Warn("Shutdown server error|%+v", err)
	} else {
		log.Info("Shutdown ok")
	}
}
