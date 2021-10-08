package models

const (
	HeaderTrace          = "X-Trace-Id"      // 将请求标识(traceID)放在 header 中
	HeaderAcceptLanguage = "Accept-Language" // 浏览器语言

	LogKeyTraceID = "traceID" // 日志中打印当前请求标识

	GinKeyCtx = "ctx"
	GinKeyT   = "t"
)
