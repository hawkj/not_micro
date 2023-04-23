package kafkajob

import (
	"context"
	"github.com/hawkj/not_micro/pkg/common"
)

type KafakJobFun func(ctx context.Context, g *common.Global, msg []byte)
