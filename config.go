package main

import (
	"fmt"
	"os"
)

type Config struct {
	MailSenderAdress    string
	MailSenderPassword  string
	MailReceiverAddress string
}

var config *Config

func loadConfig() error {
	mailSenderAdress := os.Getenv("MAIL_SENDER_ADDRESS")
	if mailSenderAdress == "" {
		return fmt.Errorf("Env var MAIL_SENDER_ADDRESS is required")
	}

	mailSenderPassword := os.Getenv("MAIL_SENDER_PASSWORD")
	if mailSenderPassword == "" {
		return fmt.Errorf("Env var MAIL_SENDER_PASSWORD is required")
	}

	mailReceiverAdress := os.Getenv("MAIL_RECEIVER_ADDRESS")
	if mailReceiverAdress == "" {
		return fmt.Errorf("Env var MAIL_RECEIVER_ADDRESS is required")
	}

	config = &Config{
		MailSenderAdress:    mailSenderAdress,
		MailSenderPassword:  mailSenderPassword,
		MailReceiverAddress: mailReceiverAdress,
	}

	return nil
}
