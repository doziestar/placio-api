package emails

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/resendlabs/resend-go"
)

type EmailType string

const (
	WelcomeEmailType       EmailType = "welcome"
	PasswordResetEmailType EmailType = "password_reset"
)

type EmailTemplateData struct {
	UserName     string
	EmailContent EmailContent
	ButtonURL    string
	ButtonLabel  string
}

type EmailContent struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Title   string `json:"title"`
}

type EmailContents map[EmailType]EmailContent

func SendCustomEmail(userName, userEmail string, emailType EmailType) error {
	apiKey := "re_dELkzomS_N7afBBFuy7uy3j7B9tyv3aqf"
	client := resend.NewClient(apiKey)
	subject, html, err := getEmailTemplate(userName, emailType)
	if err != nil {
		return err
	}

	params := &resend.SendEmailRequest{
		From:    "Placio Ltd <support@updates.placio.io>",
		To:      []string{userEmail},
		Html:    html,
		Subject: subject,
		//Cc:      []string{"dozie@placio.io"},
		//Bcc:     []string{"bcc@example.com"},
		ReplyTo: "dozie@placio.io",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		return err
	}
	fmt.Println(sent.Id)
	return nil
}

func getEmailTemplate(userName string, emailType EmailType) (subject, body string, err error) {
	// load the email contents from content.json
	contents, err := loadEmailContents()
	if err != nil {
		return "", "", err
	}
	emailContent, ok := contents[emailType]
	if !ok {
		return "", "", fmt.Errorf("no email content for type: %s", emailType)
	}
	subject = emailContent.Title

	// Get the absolute path to the template.html file
	_, currentFilePath, _, _ := runtime.Caller(0)
	templateFilePath := filepath.Join(filepath.Dir(currentFilePath), "template.html")

	// load the email template from template.html
	templateFile, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		return "", "", err
	}
	tmpl, err := template.New("email").Parse(string(templateFile))
	if err != nil {
		return "", "", err
	}

	// execute the template with the EmailTemplateData as the data
	var tpl bytes.Buffer
	data := EmailTemplateData{
		UserName:     userName,
		EmailContent: emailContent,
		ButtonURL:    "https://placio.io", // substitute with your actual button URL
		ButtonLabel:  "Click here to login",
	}
	if err := tmpl.Execute(&tpl, data); err != nil {
		return "", "", err
	}
	body = tpl.String()

	return subject, body, nil
}

func loadEmailContents() (EmailContents, error) {
	contents := make(EmailContents)

	// Get the absolute path to the content.json file
	_, currentFilePath, _, _ := runtime.Caller(0)
	contentFilePath := filepath.Join(filepath.Dir(currentFilePath), "content.json")

	contentFile, err := ioutil.ReadFile(contentFilePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(contentFile, &contents)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
