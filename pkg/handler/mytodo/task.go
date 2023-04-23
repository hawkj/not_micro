package handlermytodo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hawkj/not_micro/pkg/common"
	requestcontext "github.com/hawkj/not_micro/pkg/context"
	"github.com/hawkj/not_micro/pkg/kafka"
	mytodoservice "github.com/hawkj/not_micro/pkg/service/my_todo"
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
