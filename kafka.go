package kafka

import (
	"strings"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

const (
	// OffsetNewest stands for the log head offset, i.e. the offset that will be
	// assigned to the next message that will be produced to the partition. You
	// can send this to a client's GetOffset method to get this offset, or when
	// calling ConsumePartition to start consuming new messages.
	OffsetNewest int64 = -1
	// OffsetOldest stands for the oldest offset available on the broker for a
	// partition. You can send this to a client's GetOffset method to get this
	// offset, or when calling ConsumePartition to start consuming from the
	// oldest offset that is still available on the broker.
	OffsetOldest int64 = -2
)

// NewConsumer create a new cluster consumer
func NewConsumer(config *Config, topics []string) (*cluster.Consumer, error) {
	return cluster.NewConsumer(splitByComma(config.Brokers), config.GroupID, topics, clusterConfig())
}

// NewAsyncProducer create a async producer
func NewAsyncProducer(config *Config) (sarama.AsyncProducer, error) {
	return sarama.NewAsyncProducer(splitByComma(config.Brokers), asyncProducerConfig())
}

// NewSyncProducer create a sync producer
func NewSyncProducer(config *Config) (sarama.SyncProducer, error) {
	return sarama.NewSyncProducer(splitByComma(config.Brokers), syncProducerConfig(time.Duration(config.Timeout)))
}

func clusterConfig() *cluster.Config {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = OffsetOldest
	return config
}

func asyncProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond
	return config
}

func syncProducerConfig(timeout time.Duration) *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = timeout
	return config
}

func splitByComma(source string) []string {
	return strings.Split(source, ",")
}
