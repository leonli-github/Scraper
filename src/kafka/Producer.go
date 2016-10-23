package kafka

import (
	c "configuration"

	"github.com/Shopify/sarama"

	"log"
	"os"
	"strings"
	"debugger"
	"time"
)

func SendMessage() {
	if c.Config.SaramaLogging {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	brokerList := c.Config.KafkaBrokers
	debugger.Logf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	producer := initProducer(brokerList)

	send(producer, "important", time.Now().String())

	defer func() {
		if err := producer.Close(); err != nil {
			debugger.Logf("Failed to close producer", err)
		}
	}()
}

func initProducer(brokerList []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message

	// On the broker side, you may want to change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return producer
}

func send(producer sarama.SyncProducer, topic string, value string) {
	// We are not setting a message key, which means that all messages will
	// be distributed randomly over the different partitions.
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	})

	if err != nil {
		debugger.Logf("Failed to store your data:, %s", err)
	} else {
		// The tuple (topic, partition, offset) can be used as a unique identifier
		// for a message in a Kafka cluster.
		debugger.Logf("Your data is stored with unique identifier %s/%d/%d", topic, partition, offset)
	}
}
