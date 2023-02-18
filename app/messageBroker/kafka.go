package messageBroker

import "fmt"

// type kafkaBroker struct {
// 	producer *kafka.Producer
// }

type Publisher interface {
	Publish(topic string, message []byte) error
}

// func NewKafkaClient(host string) *kafkaBroker {
// 	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": host})
// 	if err != nil {
// 		log.Panicf("Error creating Kafka client: %+v", err)
// 	}

// 	return &kafkaBroker{producer}
// }

// func (kb *kafkaBroker) Publish(topic string, message []byte) error {
// 	producerErr := kb.producer.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 		Value:          message,
// 	}, nil)

// 	if producerErr != nil {
// 		fmt.Println("unable to enqueue message ", message)
// 	}

// 	event := <-kb.producer.Events()

// 	outMessage := event.(*kafka.Message)
// 	if outMessage.TopicPartition.Error != nil {
// 		producerErr = outMessage.TopicPartition.Error
// 	} else {
// 		fmt.Printf("Message delivery completed successfully: %s", message)
// 	}

// 	return producerErr
// }

type fakeBroker struct{}

func NewFakeBroker() *fakeBroker {
	return &fakeBroker{}
}

func (f *fakeBroker) Publish(topic string, message []byte) error {
	fmt.Printf("Publishing message: %s on topic: %s\n", string(message), topic)
	return nil
}
