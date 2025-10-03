# Email Sender Cloud Function

A Google Cloud Function that sends emails using Gmail's SMTP server. This function can be deployed to Google Cloud Functions and triggered via HTTP requests.

## Prerequisites

1. Google Cloud account with billing enabled
2. Google Cloud SDK installed and configured
3. Go 1.21 or later
4. A Gmail account (or Google Workspace account) with "Less secure app access" enabled or an App Password

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/email-sender.git
   cd email-sender
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the project root with the following variables:
   ```
   EMAIL_USER=your-email@gmail.com
   EMAIL_PASS=your-app-password
   TO_EMAIL=recipient@example.com
   ```

   > **Note:** For Gmail, you'll need to use an "App Password" if you have 2FA enabled. You can generate one at: https://myaccount.google.com/apppasswords

## Local Development

To test the function locally:

1. Run the function locally using the Functions Framework:
   ```bash
   go run github.com/GoogleCloudPlatform/functions-framework-go/cmd/functions-framework --target=SendEmailHTTP
   ```

2. Send a test request:
   ```bash
   curl -X POST http://localhost:8080 \
     -H "Content-Type: application/json" \
     -d '{"sender_email":"test@example.com", "name":"Test User", "body":"This is a test email"}'
   ```

## Deployment

### Prerequisites for Deployment

1. Create a new Google Cloud Project or use an existing one
2. Enable the Cloud Functions API
3. Create a service account with the following roles:
   - Cloud Functions Admin
   - Service Account User
   - Storage Admin
4. Generate a JSON key for the service account and add it as a GitHub secret named `GCP_SA_KEY`

### GitHub Secrets

Add the following secrets to your GitHub repository:
- `GCP_PROJECT_ID`: Your Google Cloud Project ID
- `GCP_SA_KEY`: The JSON key of your service account
- `EMAIL_USER`: Your Gmail address
- `EMAIL_PASS`: Your Gmail App Password
- `TO_EMAIL`: The recipient email address

### Deploy

Push to the `main` branch to trigger the GitHub Actions workflow, or manually trigger it from the Actions tab.

## API Usage

### Request

```http
POST /
Content-Type: application/json

{
  "sender_email": "sender@example.com",
  "name": "John Doe",
  "body": "This is the email content"
}
```

### Response

Success (200 OK):
```json
{
  "message": "Email sent successfully"
}
```

Error (4xx/5xx):
```json
{
  "error": "Error message"
}
```

## Security Considerations

1. **App Passwords**: Use App Passwords instead of your main Gmail password
2. **CORS**: The function allows requests from any origin (`*`). Restrict this in production.
3. **Environment Variables**: Never commit sensitive information to version control. Use GitHub Secrets for deployment.
4. **Rate Limiting**: Consider implementing rate limiting to prevent abuse.

## License

MIT
