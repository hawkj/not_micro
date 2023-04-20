package kafkajob

import (
	"context"
	"fmt"
	"gitee.com/lichuan2022/my-todo/pkg/common"
)

func Test(ctx context.Context, g *common.Global, msg []byte) {
	fmt.Println(string(msg))
}
