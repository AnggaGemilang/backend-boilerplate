package commontransmittergateway

type SimpleTransmitterInterface interface {
	SendMessage(payload []byte) error
}

type GroupTransmitterInterface interface {
	SendMessageToGroup(group string, payload string) error
}

type TopicTransmitterInterface interface {
	Publish(topic string, payload []byte) error
}
