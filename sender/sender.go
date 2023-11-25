package sender

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v3"
)

type MailGunParams struct {
	From    string
	Subject string
	Text    string
	Tags    []string
	To      string
}

// Sends message using Mailgun API.
func SendMailGunMessageV3(domain, apiKey string, params *MailGunParams) (string, string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	message := mg.NewMessage(
		params.From,
		params.Subject,
		params.Text,
	)
	for _, tag := range params.Tags {
		err := message.AddTag(tag)
		if err != nil {
			fmt.Println(err)
			log.Fatalf("Could not add a tag <%s>. Passed params: %+v", tag, params)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// Leaving "send" here to test out funcionality of mailgun module.
	responseMessage, id, err := mg.Send(ctx, message)
	return responseMessage, id, err
}
