package services

import (
	"encoding/json"
	"time"

	"github.com/shankar524/company-management-service/app/messageBroker"
)

type JSON map[string]interface{}

type Broadcaster interface {
	Broadcast(topic string, message JSON) (err error)
}

type messageService struct {
	Broker messageBroker.Publisher
}

func NewMessageService(broker messageBroker.Publisher) *messageService {
	return &messageService{broker}
}

func (ms *messageService) Broadcast(topic string, message JSON) (err error) {

	message["sent_at"] = time.Now().Unix()
	payload, err := json.Marshal(message)
	if err != nil {
		return
	}

	err = ms.Broker.Publish(topic, payload)
	return
}
