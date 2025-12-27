# Mail Sender

A concurrent email sending application written in Go that reads recipient data from a CSV file and sends personalized emails using HTML templates.

## Features

- 📧 **Concurrent Email Sending**: Uses a worker pool pattern to send emails concurrently for improved performance
- 📄 **CSV-based Recipients**: Reads recipient data (name and email) from CSV files
- 🎨 **Template Support**: Supports HTML email templates with dynamic content
- 🐳 **Docker Support**: Fully containerized with Docker and Docker Compose
- ⚙️ **Environment-based Configuration**: Uses environment variables for flexible configuration
- 🔒 **Error Handling**: Comprehensive error handling with detailed logging

## Project Structure

```
mailSender/
├── main.go                    # Application entry point
├── internal/
│   ├── app/
│   │   └── app.go            # Main application logic
│   ├── config/
│   │   └── config.go         # Configuration management
│   ├── file_reader/
│   │   └── csv_reader.go     # CSV file reading logic
│   └── mail_sender/
│       └── mail_sender.go    # Email template data structures
├── letter.tmpl                # Email template file
├── test.csv                   # Sample CSV file with recipients
├── Dockerfile                 # Docker build configuration
├── docker-compose.yml         # Docker Compose configuration
├── .dockerignore             # Docker ignore patterns
└── README.md                  # This file
```

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (optional, for containerized deployment)
- SMTP server credentials (Gmail, SendGrid, etc.)

## Installation

### Local Development

1. Clone the repository:
```bash
git clone <repository-url>
cd mailSender
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o mailSender .
```

## Configuration

Create a `.env` file in the root directory with the following variables:

```env
# SMTP Configuration
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
FROM=your-email@gmail.com
PASSWORD=your-app-password

# File Paths
EMAILS_PATH=test.csv
TEMPLATES_PATH=letter.tmpl
```

### Gmail Setup

If using Gmail, you'll need to:
1. Enable 2-Step Verification
2. Generate an App Password: [Google App Passwords](https://myaccount.google.com/apppasswords)
3. Use the app password in the `PASSWORD` environment variable

### Other SMTP Providers

The application works with any SMTP server. Update the `SMTP_HOST` and `SMTP_PORT` accordingly:
- **Gmail**: `smtp.gmail.com:587`
- **SendGrid**: `smtp.sendgrid.net:587`
- **Outlook**: `smtp-mail.outlook.com:587`
- **Custom SMTP**: Use your provider's SMTP settings

## Usage

### Local Execution

1. Ensure your `.env` file is configured correctly
2. Prepare your CSV file with the following format:
```csv
name,email
John Doe,john@example.com
Jane Smith,jane@example.com
```

3. Prepare your email template (`letter.tmpl`):
```
Subject: Your Subject Here

Hello {{.Name}}!

Your email content here.
```

4. Run the application:
```bash
./mailSender
```

### Using Docker

#### Option 1: Docker Build and Run

1. Build the Docker image:
```bash
docker build -t mail-sender:latest .
```

2. Run the container:
```bash
docker run --env-file .env mail-sender:latest
```

#### Option 2: Docker Compose

1. Ensure your `.env` file is in the root directory

2. Build and run:
```bash
docker-compose up --build
```

3. Run in detached mode:
```bash
docker-compose up -d
```

4. View logs:
```bash
docker-compose logs -f
```

5. Stop the container:
```bash
docker-compose down
```

## Email Template

The email template uses Go's `html/template` package. Available template variables:

- `{{.Name}}` - Recipient's name from CSV

Example template (`letter.tmpl`):
```
Subject: Welcome!

Hello {{.Name}},

Thank you for joining us!

Best regards,
The Team
```

## CSV Format

The CSV file should have the following structure:

```csv
name,email
John Doe,john@example.com
Jane Smith,jane@example.com
```

- First row is treated as header and will be skipped
- `name` column: Recipient's name (can be empty)
- `email` column: Recipient's email address (required)

## Concurrency

The application uses a worker pool pattern with 5 concurrent workers by default. This can be adjusted in `internal/app/app.go`:

```go
const numWorkers = 5 // Adjust this value
```

## Logging

The application provides detailed logging:
- Worker ID for each email operation
- Success/failure status for each email
- Error messages with context

Example log output:
```
[Worker 0] Successfully sent email to John Doe (john@example.com)
[Worker 1] Successfully sent email to Jane Smith (jane@example.com)
[Worker 2] Failed to send email to Bob (bob@example.com): failed to send mail: ...
```

## Error Handling

The application handles various error scenarios:
- Missing or invalid environment variables
- CSV file reading errors
- Template parsing errors
- SMTP connection errors
- Individual email sending failures (logged but don't stop the process)

## Development

### Running Tests

```bash
go test ./...
```

### Code Structure

- **main.go**: Entry point, initializes and runs the application
- **internal/app/app.go**: Core application logic with worker pool implementation
- **internal/config/**: Configuration management from environment variables
- **internal/file_reader/**: CSV parsing and validation
- **internal/mail_sender/**: Email template data structures

## Troubleshooting

### Common Issues

1. **"SMTP_HOST environment variable is not set"**
   - Ensure your `.env` file exists and contains all required variables

2. **"failed to send mail: authentication failed"**
   - Verify your SMTP credentials
   - For Gmail, ensure you're using an App Password, not your regular password

3. **"failed to read CSV file"**
   - Check that the file path in `EMAILS_PATH` is correct
   - Verify the CSV file format matches the expected structure

4. **"failed to parse template"**
   - Ensure the template file exists at the path specified in `TEMPLATES_PATH`
   - Check for syntax errors in the template file

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
