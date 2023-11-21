package sender

import (
	"fmt"

	"github.com/mailgun/mailgun-go"
)

type MailGunParams struct {
	from    string
	subject string
	text    string
	tags    []string
	to      []string
}

func SendMailGunMessageV3(domain, apiKey string, params MailGunParams) (string, string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		params.from,
		params.subject,
		params.text,
	)
	for _, tag := range params.tags {
		m.AddTag(tag)
	}
	for _, recipient := range params.to {
		m.AddRecipient(recipient)
	}

	msg, id, err := mg.Send(m)
	fmt.Printf("%s\n%s\n%s", msg, id, err)
	return msg, id, err
}
