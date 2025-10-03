package email_sender

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("SendEmailHTTP", SendEmailHTTP)
}

// Message represents the JSON request body
type Message struct {
	SenderEmail string `json:"sender_email"`
	Name       string `json:"name"`
	Body       string `json:"body"`
	ToEmail    string `json:"to_email"`
}

// SendEmailHTTP is an HTTP Cloud Function
type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SendEmailHTTP(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow POST requests
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{Error: "Method not allowed"})
		return
	}

	// Parse request body
	var req Message
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid request body"})
		return
	}

	// Validate required fields
	if req.SenderEmail == "" || req.Name == "" || req.Body == "" || req.ToEmail == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Missing required fields: sender_email, name, body, and to_email are required"})
		return
	}

	// Basic email format validation
	if !strings.Contains(req.ToEmail, "@") || !strings.Contains(req.ToEmail, ".") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid email format for to_email"})
		return
	}

	// Send email
	_, err := SendEmail(r.Context(), EmailRequest(req))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Email sent successfully"})
}
