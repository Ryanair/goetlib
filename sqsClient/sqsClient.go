package sqsClient

import (
	"errors"
	"github.com/Ryanair/goetlib/logger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/sqs"
	"strings"
)

type SqsPublisher struct {
	client   *sqs.SQS
	queueUrl *string
	Publisher
}

type Publisher interface {
	publish(body string) error
}

func NewSqsPublisher(client *sqs.SQS, queueUrl *string) *SqsPublisher {
	return &SqsPublisher{
		client:   client,
		queueUrl: queueUrl,
	}
}

func (scheduler *SqsPublisher) Publish(body string) error {
	if scheduler.queueUrl == nil {
		return errors.New("sending message to default queue which is undefined")
	}

	input := &sqs.SendMessageInput{
		MessageBody: &body,
		QueueUrl:    scheduler.queueUrl,
	}

	logger.Debug("Sending message to SQS: %+v", input)
	_, err := scheduler.client.SendMessage(input)
	if err != nil {
		logger.Error("Error sending message %+v caused error %+v", input, err)
		return err
	}
	return nil
}

func ValidateEvent(event events.SQSMessage) (string, error) {
	if len(strings.TrimSpace(event.Body)) == 0 {
		return "", errors.New("event body is mandatory")
	}

	return event.Body, nil
}
