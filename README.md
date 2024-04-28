# Mailer Module

The mailer module provides a convenient way to send emails using Go. It includes functionality for validating email data and sending emails with SMTP authentication.

## Features
- Simple and intuitive API for sending emails.
- Support for email authentication configuration.
- Templating support for email content customization.

## Installation

To use the mailer module in your Go project, you can simply import it using:

```go
import "github.com/vidarlx/mailer"
```

## Usage

### Creating a Mailer

To create a new mailer instance, use the NewMailer function:

```go
authConfig := mailer.EmailAuthConfig{
    Host:     "smtp.example.com",
    Port:     587,
    Username: "user@example.com",
    Password: "password",
    From:     "user@example.com",
}
templateDir := "templates"

mailService := mailer.NewMailer(authConfig, templateDir)
```

### Sending an Email
To send an email, create an Email struct with the necessary data and pass it to the SendEmail method of the mailer instance:

```go
email := &mailer.Email{
    Recipients:       []string{"recipient@example.com"},
    Subject:          "Test Subject",
    Body:             "Test Body",
    TemplateFileName: "test_template.tmpl",
    Data:             nil, // Assuming no data for simplicity
}

err := mailService.SendEmail(email)
if err != nil {
    // Handle error
}
```

### Templating Support

You can use templates to customize the content of your emails. Simply provide the template file name and data to be parsed with the template:

```go
email := &mailer.Email{
    Recipients:       []string{"recipient@example.com"},
    Subject:          "Test Subject",
    TemplateFileName: "template.tmpl",
    Data: map[string]interface{}{
        "Name": "John Doe",
    },
}
```

### Validation

The mailer module provides validation for email data to ensure required fields are present:

```go
email := &mailer.Email{
    Recipients:       []string{},
    Subject:          "Test Subject",
    Body:             "Test Body",
    TemplateFileName: "test_template.tmpl",
    Data:             nil,
}

err := mailer.ValidateEmail(email)
if err != nil {
    // Handle validation error
}
```

License

This module is released under the MIT License. Feel free to use and modify it according to your needs. I don't care.
