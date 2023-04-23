package kafkajob

import (
	"context"
	"fmt"
	"github.com/hawkj/not_micro/pkg/common"
)

func Test(ctx context.Context, g *common.Global, msg []byte) {
	fmt.Println(string(msg))
}
