package ntfy

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/nikoksr/notify/v2"
)

func sendRequest(client *http.Client, req *http.Request) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := asNotifyError(resp); err != nil {
		return err
	}

	return nil
}

func (s *Service) sendTextMessage(ctx context.Context, topic string, conf SendConfig) error {
	payload := &sendMessageRequest{
		Topic:       topic,
		Title:       conf.subject,
		Message:     conf.message,
		Tags:        conf.tags,
		Priority:    conf.priority,
		ClickAction: conf.clickAction,
		Markdown:    conf.parseMode == ModeMarkdown,
		Delay:       conf.delay,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.apiBaseURL, bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.token)

	return sendRequest(s.client, req)
}

func (s *Service) sendFile(ctx context.Context, topic string, attachment notify.Attachment) error {
	// Append topic to base URL, e.g. https://ntfy.sh/my_topic
	endpoint := s.apiBaseURL + topic

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, endpoint, attachment)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+s.token)
	// req.Header.Set("X-Title", attachment.Name())

	return sendRequest(s.client, req)
}

func (s *Service) sendFileAttachments(ctx context.Context, topic string, conf SendConfig) error {
	for _, attachment := range conf.attachments {
		if err := s.sendFile(ctx, topic, attachment); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) sendToTopic(ctx context.Context, topic string, conf SendConfig) error {
	if err := s.sendTextMessage(ctx, topic, conf); err != nil {
		return err
	}

	return s.sendFileAttachments(ctx, topic, conf)
}

// Send sends a message to all topics that are configured to receive messages. It returns an error if the message could
// not be sent.
func (s *Service) Send(ctx context.Context, subject, message string, opts ...notify.SendOption) error {
	conf := SendConfig{
		subject:     subject,
		message:     message,
		parseMode:   s.parseMode,
		priority:    s.priority,
		tags:        s.tags,
		delay:       s.delay,
		clickAction: s.clickAction,
	}

	for _, opt := range opts {
		opt(&conf)
	}

	for _, topic := range s.topics {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := s.sendToTopic(ctx, topic, conf); err != nil {
			return &notify.SendNotificationError{
				Recipient: topic,
				Cause:     err,
			}
		}
	}

	return nil
}