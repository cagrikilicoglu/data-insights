package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

type SMTPEmailService struct {
	SMTPHost     string
	SMTPPort     string
	SenderEmail  string
	SenderPasswd string
}

func NewSMTPEmailService(host, port, email, passwd string) *SMTPEmailService {
	return &SMTPEmailService{
		SMTPHost:     host,
		SMTPPort:     port,
		SenderEmail:  email,
		SenderPasswd: passwd,
	}
}

// SendEmail sends an email to the specified recipient with the provided subject and body.
// It connects to the SMTP server, secures the connection with TLS, authenticates, and sends the email.
func (s *SMTPEmailService) SendEmail(to string, subject string, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", s.SenderEmail, s.SenderPasswd, s.SMTPHost)

	// Connect to the SMTP server.
	client, err := smtp.Dial(s.SMTPHost + ":" + s.SMTPPort)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer client.Quit()

	// Upgrade the connection to TLS.
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // TODO Change in production
		ServerName:         s.SMTPHost,
	}

	if err = client.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("failed to start TLS: %v", err)
	}

	// Authenticate.
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate to SMTP server: %v", err)
	}

	// Set the sender and recipient.
	if err = client.Mail(s.SenderEmail); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %v", err)
	}

	// Get the writer for the email data.
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to send data: %v", err)
	}

	// Format the email with the correct MIME type for HTML.
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s",
		s.SenderEmail, to, subject, body)

	// Write the email content.
	if _, err = w.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write message: %v", err)
	}

	// Close the writer to finish the email.
	if err = w.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %v", err)
	}

	return nil
}
