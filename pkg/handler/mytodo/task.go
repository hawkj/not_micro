package handlermytodo

import (
	"fmt"
	"gitee.com/lichuan2022/my-todo/pkg/common"
	requestcontext "gitee.com/lichuan2022/my-todo/pkg/context"
	"gitee.com/lichuan2022/my-todo/pkg/kafka"
	mytodoservice "gitee.com/lichuan2022/my-todo/pkg/service/my_todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TaskCreate(c *requestcontext.CommonContext) {
	taskId, err := mytodoservice.TaskCreate(c)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.GinContext.JSON(http.StatusOK, map[string]string{
		"task_id": taskId,
	})
}

func PostMsg(c *requestcontext.CommonContext) {
	p, err := kafka.NewProducer([]string{c.Config.Kafka.Host}, common.KafkaTopicTest)
	if err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("test-msg-%d", time.Now().Unix())
	if err := p.SendMessage(message); err != nil {
		c.GinContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.GinContext.JSON(http.StatusOK, map[string]string{
		"msg": "send msg done",
	})
}
