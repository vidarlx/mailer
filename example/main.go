package main

import (
	"log"

	"github.com/vidarlx/mailer"
)

func main() {
	// Define email authentication configuration
	auth := mailer.EmailAuthConfig{
		Host:     "smtp.mandrillapp.com",
		Port:     587,
		Username: "anyname",
		Password: "so.secret!",
		From:     "me@somesender.hi",
	}

	// Create a new mailer instance
	mailService := mailer.NewMailer(auth, "./templates")

	// Send newsletter
	Newsletter(mailService)
}

func Newsletter(mail mailer.Mailer) {
	articles := []struct {
		Title   string
		Content string
	}{
		{
			Title:   "The Joy of Waiting in Line",
			Content: "Ah, the thrill of standing in line! Whether it's at the DMV or the grocery store, there's nothing quite like the excitement of wasting your precious time while waiting for service. Some say it's a form of meditation, but we all know it's just a delightful inconvenience.",
		},
		{
			Title:   "The Art of Small Talk",
			Content: "Who doesn't love engaging in mind-numbing small talk? From discussing the weather for the umpteenth time to pretending to care about your neighbor's cousin's cat, there's no shortage of riveting conversation topics. After all, nothing says 'I'm genuinely interested' like a forced smile and a polite nod.",
		},
	}

	// Create email with newsletter data
	email := &mailer.Email{
		Recipients:       []string{"bob@hello.world", "cindy@hello.world"},
		Subject:          "Weekly Newsletter",
		TemplateFileName: "newsletter.html",
		Data: map[string]interface{}{
			"Articles": articles,
		},
	}

	// Send email
	err := mail.SendEmail(email)
	if err != nil {
		log.Fatal(err)
	}
}
