# Data Insights

**Data Insights** is a tool designed to analyze web analytics data, aggregate key performance metrics, and generate insights using a Large Language Model (LLM). The insights are then sent via email, formatted in a user-friendly report.

## Features

- **Data Aggregation**: Processes JSON data files to calculate various metrics such as engagement rates, bounce rates, session durations, and more.
- **LLM Integration**: Generates AI-driven insights using OpenAI's API.
- **Email Reporting**: Sends a detailed report of the analysis via email.
- **Configurable Thresholds**: Allows customization of data processing thresholds for more tailored analysis.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/cagrikilicoglu/data-insights.git
   cd data-insights

2. Add data files under files folder
3. Install dependencies:
   go mod tidy

## Usage

1. Set Up Environment Variables: Create a .env file in the root directory with the following variables:

```bash
FILE_DIR=your-data-directory
OPENAI_API_KEY=your-openai-api-key
EMAIL_FROM=your-email@example.com
EMAIL_FROM_PASS=your-email-password
EMAIL_TO=recipient-email@example.com
RECIPIENT_NAME=Recipient Name
SMTP_HOST=smtp.your-email-provider.com
SMTP_PORT=your-smtp-port
```

2. Add data files to files folder
3. Run the Application:
    
```bash
  go run cmd/main.go
```

## Configuration

- Threshold Values: The threshold for the minimum number of data points can be adjusted in the metrics package (thresholdDataPointNumber constant).
- Email Templates: The HTML template for the email report is located in the templates directory.

## Project Structure

```bash
data-insights/
├── cmd/
│   ├── main.go               # Entry point for the application
│   └── bootstrap.go          # Main functionality of app
├── pkg/
│   ├── service.go            # Main business logic and service handling
├── kit/
│   ├── ai/
│   │   ├── client.go         # OpenAI client and interaction logic
│   │   ├── const.go          # AI related constants
│   │   ├── model.go          # AI related models
│   │   └── prompt.go         # Prompt creation logic
│   ├── email/
│   │   ├── smtp.go           # SMTP email service for sending reports
│   │   ├── const.go          # Email related constants
│   │   ├── service.go        # Email service interface
│   │   └── renderer.go       # Email template rendering logic
│   ├── file/
│   │   ├── json.go           # Data parsing logic from json files
│   │   └── util.go           # File handling utilities for reading data files
│   ├── metrics/
│   │   ├── aggregated.go     # Logic for aggregating metrics by breakdowns
│   │   ├── overall.go        # Logic for calculating overall metrics
│   │   ├── util.go           # Metric utilities
│   │   └── ui.go             # Interface for gathering important metrics
│   └── common/
│   │   ├── consts.go         # Common constants
│   │   ├── sort.go           # Metric sorting
│       └── model.go          # Common models used across the project
├── templates/
│   └── email_template.html   # HTML template for the email report
├── files/                    # Data files to be analyzed
├── go.mod                    # Example environment variables file
├── .env                      # Example environment variables file
├── LICENSE                   # License file
└── README.md                 # Project documentation
```

## Contributing

Feel free to open issues or submit pull requests if you'd like to contribute to this project.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

For any inquiries, please reach out to mcagrikilicoglu@gmail.com.
