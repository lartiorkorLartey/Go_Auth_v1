package utils

import (
	"bytes"
	"html/template"
	"os"
    "fmt"

	"gopkg.in/gomail.v2"
)

type FeatureRequest struct {
    FeatureName        string `json:"feature_name" binding:"required"`
    FeatureDescription string `json:"feature_description" binding:"required"`
}

func SendFeatureRequestEmail( request FeatureRequest) error {
    from := os.Getenv("EMAIL_ADDRESS")
    password := os.Getenv("EMAIL_PASSWORD")

    tmpl, err := template.ParseFiles("templates/feature_request_email.html")
    if err != nil {
        return err
    }

    var body bytes.Buffer
    if err := tmpl.Execute(&body, request); err != nil {
        return err
    }

    mailer := gomail.NewMessage()
    mailer.SetHeader("From", from)
    mailer.SetHeader("To", from)
    mailer.SetHeader("Subject", "New Feature Request")
    mailer.SetBody("text/html", body.String())

    dialer := gomail.NewDialer("smtp.gmail.com", 587, from, password)
    dialer.TLSConfig = nil

    if err := dialer.DialAndSend(mailer); err != nil {
        fmt.Println(err)
        return err
    }

    return nil
}

