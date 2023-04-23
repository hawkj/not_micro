package requestcontext

import (
	"github.com/gin-gonic/gin"
	"github.com/hawkj/not_micro/config"
	"github.com/hawkj/not_micro/pkg/common"
)

type CommonContext struct {
	GinContext *gin.Context
	Global     *common.Global
	Uid        string
	Config     *config.Config
}
