package email

import (
	"bytes"
	"data-insights/kit/common"
	"html/template"
)

type Renderer interface {
	Render(emailData common.EmailData) (string, error)
}

type EmailRenderer struct {
	templatePath string
}

func NewRenderer(templatePath string) *EmailRenderer {
	return &EmailRenderer{templatePath: templatePath}
}

// Render parses the email template, executes the parsed template with the provided email data and
// returns the resulting email body as a string.
func (r *EmailRenderer) Render(emailData common.EmailData) (string, error) {
	tmpl, err := template.ParseFiles(r.templatePath)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, emailData); err != nil {
		return "", err
	}

	return body.String(), nil
}
