package handlermytodo

import (
	"gitee.com/lichuan2022/my-todo/pkg/context"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *requestcontext.CommonContext) {
	c.GinContext.JSON(http.StatusOK, gin.H{"message": "pong"})
}
