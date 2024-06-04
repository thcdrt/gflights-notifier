package main

import (
	"fmt"

	"github.com/pkg/errors"
	gomail "gopkg.in/mail.v2"
)

type AvailableFlight struct {
	Date        string
	Price       string
	TravelClass string
}

func sendAvailabilitiesByMail(availabilities []string) error {
	var mailContent = "Liste des date de vols disponible avec Google Flights:\n"

	for _, availability := range availabilities {
		mailContent += fmt.Sprintf("- %s\n", availability)
	}

	mailContent += "Lien pour réserver le vol: https://www.google.com/travel/flights?tfs=CBwQARodag0IAxIJL20vMDI1ZHpicgwIAxIIL20vMGNiaGgaHWoMCAMSCC9tLzBjYmhocg0IAxIJL20vMDI1ZHpiQAFIAXABggELCP___________wGYAQE&tfu=KgIIAw&hl=fr\n"

	fmt.Println(mailContent)

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.MailSenderAdress)

	// Set E-Mail receivers
	m.SetHeader("To", config.MailReceiverAddress)

	// Set E-Mail subject
	m.SetHeader("Subject", "Disponibilités vol Nouméa -> Paris")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", mailContent)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, config.MailSenderAdress, config.MailSenderPassword)

	// Now send E-Mail
	err := d.DialAndSend(m)

	return errors.Wrap(err, "fail to send mail")
}
