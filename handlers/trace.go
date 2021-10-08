package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/youthlin/logs/pkg/kv"
	"github.com/youthlin/pub/models"
)

// Trace add a trace id to each request
func Trace(c *gin.Context) {
	traceID := GetTraceID(c)
	ctx := GetCtx(c)
	ctx = kv.Add(ctx, models.LogKeyTraceID, traceID)
	c.Set(models.GinKeyCtx, ctx)
	c.Next()
}

// GetTraceID return or generate a new traceID
func GetTraceID(c *gin.Context) string {
	trace, ok := c.Get(models.HeaderTrace)
	if ok {
		return trace.(string)
	}
	traceID := c.GetHeader(models.HeaderTrace)
	if traceID == "" {
		traceID = uuid.NewString()
	}
	// set to gin.Context and client response header
	c.Set(models.HeaderTrace, traceID)    // 写入 Context 下次 GetTraceID 就能拿到
	c.Header(models.HeaderTrace, traceID) // 写入 Response Header 在 render 时能返回客户端
	return traceID
}

// GetCtx return the context.Context of this request
func GetCtx(c *gin.Context) context.Context {
	ctx, ok := c.Get(models.GinKeyCtx)
	if !ok {
		ctx = context.Background()
		c.Set(models.GinKeyCtx, ctx)
	}
	return ctx.(context.Context)
}
