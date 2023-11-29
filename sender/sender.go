package sender

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailGunPayload struct {
	From            string
	Subject         string
	Text            string
	TemplateVersion string // t:version
	To              string
	Tags            []string // o:tag
}

// Sends message using Mailgun API to a SINGLE recipient.
func SendMailGunMessageV4(domain, apiKey string, payload *MailGunPayload) (string, string, error) {
	mailgun := mailgun.NewMailgun(domain, apiKey)
	// Empty messages return "message not valid" from MailGun.
	if payload.Text == "" {
		payload.Text = " "
	}
	message := mailgun.NewMessage(payload.From, payload.Subject, payload.Text, payload.To)
	// Set correct template version (language).
	message.SetTemplateVersion(payload.TemplateVersion)
	// Add tags.
	for _, tag := range payload.Tags {
		err := message.AddTag(tag)
		if err != nil {
			fmt.Println(err)
			log.Fatalf("Could not add a tag <%s>. Passed payload: %+v", tag, payload)
		}
	}
	// Send message with timeout to ensure resources are realeased after message was sent.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	responseMessage, id, err := mailgun.Send(ctx, message)
	return responseMessage, id, err
}
