package kafka

import "github.com/IBM/sarama"

type ConsumerHandler struct {
	ready     chan bool
	messageCh chan *sarama.ConsumerMessage
}

func (h *ConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

func (h *ConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg := <-claim.Messages():
			session.MarkMessage(msg, "")
			h.messageCh <- msg
		case <-session.Context().Done():
			return nil
		}
	}
}
