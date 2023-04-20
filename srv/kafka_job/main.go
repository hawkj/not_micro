package main

import (
	"context"
	"flag"
	"fmt"
	"gitee.com/lichuan2022/my-todo/config"
	"gitee.com/lichuan2022/my-todo/pkg/common"
	"gitee.com/lichuan2022/my-todo/pkg/db_mysql/my_todo"
	"gitee.com/lichuan2022/my-todo/pkg/handler/kafka_job"
	"gitee.com/lichuan2022/my-todo/pkg/redis"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

var handlerMap = map[string]kafkajob.KafakJobFun{
	"test-topic": kafkajob.Test,
}

func main() {
	env := os.Getenv("APP_ENV")
	configFile := ""
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		configFile = "/Users/testtest/StudyWorkSpace/my-todo/config/test_conf.yaml"
	} else {
		gin.SetMode(gin.DebugMode)
		configFile = "/Users/testtest/StudyWorkSpace/my-todo/config/test_conf.yaml"
	}
	// Parse command line arguments
	// -topic test-topic
	topic := flag.String("topic", "", "Name of the job to run")
	flag.Parse()

	//topicString := common.KafkaTopicTest
	//topic := &topicString
	handler, ok := handlerMap[*topic]
	if !ok {
		panic(fmt.Sprintf("Handler not found for job name: %s\n", *topic))
	}
	config := config.GetConfig(configFile)

	g := &common.Global{
		DbMyTodo: mytododb.NewDb(config.MyTodoDb),
		Redis:    redis.NewRedis(config.Redis),
	}
	ctx := context.Background()

	// Set up configuration
	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Return.Errors = true
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create new consumer
	consumer, err := sarama.NewConsumer([]string{config.Kafka.Host}, saramaConfig)
	if err != nil {
		log.Fatal("Failed to start consumer: ", err)
	}
	defer consumer.Close()

	// Determine topic to consume based on job name

	// Set up Kafka listener
	partitionConsumer, err := consumer.ConsumePartition(*topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(fmt.Sprintf("Failed to start partition consumer: %s", err.Error()))
	}
	defer partitionConsumer.Close()

	// Listen for messages
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case message := <-partitionConsumer.Messages():
			handler(ctx, g, message.Value)
		case err := <-partitionConsumer.Errors():
			log.Println("Error received while consuming partition: ", err)
		case <-signals:
			log.Println("Interrupted while listening for messages")
			return
		}
	}
}
