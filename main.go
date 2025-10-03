package email_sender

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type EmailRequest struct {
	SenderEmail string `json:"sender_email"`
	Name       string `json:"name"`
	Body       string `json:"body"`
	ToEmail    string `json:"to_email"`
}

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
}

func SendEmail(ctx context.Context, req EmailRequest) (string, error) {
	// Get environment variables
	emailUser := os.Getenv("EMAIL_USER")
	emailPass := os.Getenv("EMAIL_PASS")
	allowedEmails := strings.Split(os.Getenv("ALLOWED_EMAILS"), ",")

	// Validate recipient email is in allowed list
	isAllowed := false
	for _, email := range allowedEmails {
		if strings.TrimSpace(email) == req.ToEmail {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return "", fmt.Errorf("email not authorized to send to %s", req.ToEmail)
	}

	if emailUser == "" || emailPass == "" || req.ToEmail == "" {
		return "", fmt.Errorf("missing required environment variables")
	}

	// Create email message
	m := gomail.NewMessage()
	m.SetHeader("From", emailUser)
	m.SetHeader("To", req.ToEmail)
	m.SetHeader("Reply-To", req.SenderEmail)
	m.SetHeader("Subject", fmt.Sprintf("New message from %s", req.Name))

	// Create email body
	body := fmt.Sprintf(`
	From: %s <%s>
	Name: %s

	Message:
	%s
	`, req.Name, req.SenderEmail, req.Name, req.Body)

	m.SetBody("text/plain", strings.TrimSpace(body))

	// Send email
	d := gomail.NewDialer("smtp.gmail.com", 587, emailUser, emailPass)

	if err := d.DialAndSend(m); err != nil {
		return "", fmt.Errorf("failed to send email: %v", err)
	}

	return "Email sent successfully!", nil
}
