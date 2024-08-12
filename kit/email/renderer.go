package email

import (
	"bytes"
	"data-insights/kit/model"
	"html/template"
)

// Define the interface
type Renderer interface {
	Render(emailData model.EmailData) (string, error)
}

// Define the struct that implements the interface
type EmailRenderer struct {
	templatePath string
}

// Constructor function for EmailRenderer
func NewRenderer(templatePath string) *EmailRenderer {
	return &EmailRenderer{templatePath: templatePath}
}

// Implement the Render method for EmailRenderer
func (r *EmailRenderer) Render(emailData model.EmailData) (string, error) {
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
