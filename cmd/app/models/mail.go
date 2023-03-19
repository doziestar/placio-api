package models

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type EmailContent struct {
	Title   string `json:"title"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Button  struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	} `json:"button"`
}

var emails map[string]EmailContent

func init() {
	// Load email content from JSON file
	content, err := ioutil.ReadFile("../emails/content.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &emails)
	if err != nil {
		panic(err)
	}
}

func Send(to string, template string, data map[string]string) error {
	// Validate email address
	if !isValidEmail(to) {
		return errors.New("Invalid email address")
	}

	// Get email content
	content, ok := emails[template]
	if !ok {
		return errors.New("Email template not found")
	}

	// Create email HTML
	html, err := createEmail("template", content, data)
	if err != nil {
		return err
	}

	// Send email via Mailgun API
	values := url.Values{}
	values.Set("from", settings.sender)
	values.Set("to", to)
	values.Set("subject", content.Subject)
	values.Set("html", html)
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/messages", settings.baseURL, settings.domain), bytes.NewBufferString(values.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth("api", settings.apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Email not sent, status code: %d", resp.StatusCode)
	}

	fmt.Printf("Email sent to: %s\n", to)
	return nil
}

func createEmail(template string, content EmailContent, values map[string]string) (string, error) {
	// Load email template
	templateFile := path.Join("emails", template+".html")
	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return "", err
	}

	// Inject values into template
	email := string(templateContent)
	email = strings.ReplaceAll(email, "{{domain}}", values["domain"])

	if content.Title == "" {
		content.Title = content.Subject
	}
	if strings.Contains(content.Button.URL, "{{domain}}") {
		content.Button.URL = strings.ReplaceAll(content.Button.URL, "{{domain}}", values["domain"])
	}
	bodyLines := strings.Split(content.Body, "\n")
	if len(bodyLines) > 0 && bodyLines[0] != "" {
		bodyLines[0] = fmt.Sprintf("Hi %s,", content.Title)
	}
	body := ""
	for _, line := range bodyLines {
		body += fmt.Sprintf(`<p style="color: #7e8890; font-family: 'Source Sans Pro', helvetica, sans-serif; font-size: 15px; font-weight: normal; Margin: 0; Margin-bottom: 15px; line-height: 1.6;">%s</p>`, line)
	}
	email = strings.ReplaceAll(email, "{{title}}", content.Title)
	email = strings.ReplaceAll(email, "{{body}}", body)
	email = strings.ReplaceAll(email, "{{buttonURL}}", content.Button.URL)
	email = strings.ReplaceAll(email, "{{buttonLabel}}", content.Button.Label)
	
	email = strings.ReplaceAll(email, "{{year}}", fmt.Sprintf("%d", time.Now().Year()))

	if strings.Contains(email, "{{") {
		return "", errors.New("Email template contains invalid values")
	}

	return email, nil

}
