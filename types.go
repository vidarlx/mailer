package mailer

import "sync"

// EmailAuthConfig contains email authentication data.
//
// Example:
//
//	EmailAuthConfig{Host: "smtp.gmail.com", Port: 587, Username: "user", Password: "password", From: "me@gmail.com"}
type EmailAuthConfig struct {
	Host     string // Email provider's host
	Port     int    // Email provider's port
	Username string // Email provider's username
	Password string // Email provider's password
	From     string // Sender's email address
}

// Email contains email data.
//
// If data is provided, the template will be parsed with it.
// Otherwise, it will be parsed with the struct below.
//
//	struct { Title string; Body string }
type Email struct {
	Recipients       []string               // Email recipients
	Subject          string                 // Email subject
	Body             string                 // Email body
	TemplateFileName string                 // Template file name
	Data             map[string]interface{} // Data for template parsing
}

// MailerConfig holds email sending configuration.
type MailerConfig struct {
	Config      EmailAuthConfig
	TemplateDir string
	email       *Email
	mutex       sync.Mutex
}

type ValidateKind string

const (
	auth  ValidateKind = "auth"
	email ValidateKind = "email"
)
