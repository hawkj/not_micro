package handlermytodo

import (
	"github.com/gin-gonic/gin"
	"github.com/hawkj/not_micro/pkg/context"
	"net/http"
)

func Test(c *requestcontext.CommonContext) {
	c.GinContext.JSON(http.StatusOK, gin.H{"message": "pong"})
}
