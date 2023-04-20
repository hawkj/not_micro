package requestcontext

import (
	"gitee.com/lichuan2022/my-todo/config"
	"gitee.com/lichuan2022/my-todo/pkg/common"
	"github.com/gin-gonic/gin"
)

type CommonContext struct {
	GinContext *gin.Context
	Global     *common.Global
	Uid        string
	Config     *config.Config
}
