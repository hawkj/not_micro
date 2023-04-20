package kafkajob

import (
	"context"
	"gitee.com/lichuan2022/my-todo/pkg/common"
)

type KafakJobFun func(ctx context.Context, g *common.Global, msg []byte)
