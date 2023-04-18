package consumers

import (
	log "github.com/sirupsen/logrus"
)

func (c *KafkaConsumer) ConsumeMessages(topic string) {

	subErr := c.Consumer.Subscribe(topic, nil)
	if subErr != nil {
		log.Error("an error ocurred", subErr)
	}

	run := true
	for run {
		msg, err := c.Consumer.ReadMessage(-1)
		if err != nil {
			c.Consumer.Close()
			log.Error("an error ocurred", err)
		}

		log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	}

}
