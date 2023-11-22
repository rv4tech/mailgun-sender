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
	To      []string
}

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
	for _, recipient := range params.To {
		err := message.AddRecipient(recipient)
		if err != nil {
			fmt.Println(err)
			log.Fatalf("Could not add a recipient <%s>. Passed params: %+v", recipient, params)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// Leaving "send" here to test out funcionality of mailgun module.
	responseMessage, id, err := mg.Send(ctx, message)
	fmt.Printf("Response message: %s\n", responseMessage)
	fmt.Printf("Response ID: %s\n", id)
	fmt.Printf("Response Error: %s\n", err)
	return responseMessage, id, err
}
