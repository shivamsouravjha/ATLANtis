package utils

import (
	helpers "Atlantis/helpers/es"
	kafkaFunc "Atlantis/helpers/kafkaConsumer"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Listener(topic string) {
	c := kafkaFunc.InitConsumer(topic)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}

			helpers.EsUploader(string(ev.Key), ev.Value)
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
	}

	defer c.Close()

}
