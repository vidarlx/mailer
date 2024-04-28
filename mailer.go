package mailer

import (
	"bytes"
	"fmt"
	"net/smtp"
	"sync"
	"text/template"
)

type Mailer interface {
	Validate(kind ValidateKind) error
	SendEmail(mail *Email) error
}

func NewMailer(auth EmailAuthConfig, templateDir string) Mailer {
	return &MailerConfig{
		Config:      auth,
		TemplateDir: templateDir,
	}
}

// Validate validates the email data.
func (g *MailerConfig) Validate(kind ValidateKind) error {
	switch kind {
	case auth:
		if err := g.validateAuth(); err != nil {
			return err
		}
	case email:
		if err := g.validateEmail(); err != nil {
			return err
		}
	}
	return nil
}

func (g *MailerConfig) validateAuth() error {
	cfg := g.Config
	if cfg.Host == "" || cfg.Port == 0 || cfg.Username == "" || cfg.Password == "" || cfg.From == "" || g.TemplateDir == "" {
		return ErrInvalidConfig
	}
	return nil
}

func (g *MailerConfig) validateEmail() error {
	if len(g.email.Recipients) == 0 {
		return ErrEmptyTo
	}
	return nil
}

// SendEmail sends an email to the specified recipients with the given subject, template, and data.
func (g *MailerConfig) SendEmail(mail *Email) error {
	g.email = mail
	if err := g.Validate(email); err != nil {
		return err
	}

	if err := g.parseTemplate(); err != nil {
		return err
	}

	auth, err := g.authenticate()
	if err != nil {
		return err
	}

	err = smtp.SendMail(fmt.Sprintf("%s:%d", g.Config.Host, g.Config.Port), auth, g.Config.Username, g.email.Recipients, g.message())
	if err != nil {
		return err
	}

	return nil
}

func (g *MailerConfig) authenticate() (smtp.Auth, error) {
	if err := g.validateAuth(); err != nil {
		return nil, err
	}

	return smtp.PlainAuth("", g.Config.Username, g.Config.Password, g.Config.Host), nil
}

func (g *MailerConfig) parseTemplate() error {
	var data interface{}
	buf := new(bytes.Buffer)

	templFileName := fmt.Sprintf("%s/%s", g.TemplateDir, g.email.TemplateFileName)
	t, err := template.ParseFiles(templFileName)
	if err != nil {
		return err
	}

	if g.email.Data != nil {
		data = g.email.Data
	} else {
		data = struct {
			Title string
			Body  string
		}{
			Title: g.email.Subject,
			Body:  g.email.Body,
		}
	}

	if err = t.Execute(buf, data); err != nil {
		return err
	}

	g.email.Body = buf.String()

	return nil
}

func (g *MailerConfig) message() []byte {
	var message bytes.Buffer

	fmt.Fprintf(&message, "From: %s\r\n", g.Config.From)
	fmt.Fprintf(&message, "Subject: %s!\r\n", g.email.Subject)
	fmt.Fprintf(&message, "Content-Type: text/html; charset=UTF-8\r\n")
	fmt.Fprintf(&message, "MIME-version: 1.0\r\n")

	var wg sync.WaitGroup
	wg.Add(len(g.email.Recipients))

	for _, recipient := range g.email.Recipients {
		go func(recipient string) {
			defer wg.Done()
			g.mutex.Lock()
			defer g.mutex.Unlock()
			fmt.Fprintf(&message, "To: %s\r\n", recipient)
		}(recipient)
	}

	wg.Wait()

	message.WriteString("\r\n")
	message.WriteString(g.email.Body)

	return message.Bytes()
}
