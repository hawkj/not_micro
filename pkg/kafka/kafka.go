package kafka

import (
	"github.com/Shopify/sarama"
)

type Producer struct {
	syncProducer sarama.SyncProducer
	topic        string
}

type Consumer struct {
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
}

func NewProducer(brokerList []string, topic string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}
	return &Producer{syncProducer: producer, topic: topic}, nil
}

func (p *Producer) SendMessage(message string) error {
	defer p.syncProducer.Close()
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := p.syncProducer.SendMessage(msg)

	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) SendMessageWithKey(key string, message string) error {
	defer p.syncProducer.Close()
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}
	_, _, err := p.syncProducer.SendMessage(msg)

	if err != nil {
		return err
	}
	return nil
}

func NewConsumer(brokerList []string, topic string, partition int32) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		return nil, err
	}
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		return nil, err
	}

	return &Consumer{consumer: consumer, partitionConsumer: partitionConsumer}, nil
}

func (c *Consumer) ReadMessage() (*sarama.ConsumerMessage, error) {
	select {
	case message := <-c.partitionConsumer.Messages():
		return message, nil
	case err := <-c.partitionConsumer.Errors():
		return nil, err
	}
}
