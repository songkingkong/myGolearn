package modellearn

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestKafka(t *testing.T) {
	broker := "kafka-test-log"
	groupID := "gohangout-bilog"
	offSet := "earliest"
	topics := []string{"xxx-bi-log"}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"group.id":           groupID,
		"session.timeout.ms": 6000,
		"auto.offset.reset":  offSet})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Printf("Created Consumer %v\n", consumer)
	err = consumer.SubscribeTopics(topics, nil)
	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			t.Log("Caught signal", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				t.Log(e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					t.Log(e.Headers)
				}
			case kafka.Error:
				t.Log(e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				t.Log(e)
			}
		}
	}
	t.Log("Closing consumer")
	consumer.Close()
}
