package middleware

import (
	"context"
	"errors"
	"gitee.com/lichuan2022/my-todo/config"
	"gitee.com/lichuan2022/my-todo/pkg/common"
	"gitee.com/lichuan2022/my-todo/pkg/context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type myCommonApiHandler func(c *requestcontext.CommonContext)

func MyHandlerWrapper(handlerFunc myCommonApiHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, _ := c.Get("CommonContext")
		handlerFunc(ctx.(*requestcontext.CommonContext))
	}
}

func CommonContext(g *common.Global, conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		commonContext := &requestcontext.CommonContext{
			GinContext: c,
			Global:     g,
			Uid:        "abcdabcdabcdabcd",
			Config:     conf,
		}
		c.Set("CommonContext", commonContext)
		c.Next()
	}
}

func TimeoutMiddleware(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), duration)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()

		// 处理请求
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				c.AbortWithError(http.StatusGatewayTimeout, ctx.Err())
				return
			}
		default:
			c.Next()
		}
	}
}
