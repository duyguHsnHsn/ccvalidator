# ccvalidator

This project is created for learning purposes and to demonstrate various concepts in Go programming language,
HTTP server development, and concurrency.
---


The project implements a credit card validator microservice using the Luhn algorithm.
The microservice exposes functionality through a JSON API, allowing users to verify the validity of credit card numbers.

### Installation

1. Clone the repository:
2. Navigate to the project directory:
3. Build the project:
   ```bash
   go build
   ```
4. Run the project:
   ```bash
   ./ccvalidator
   ```
The microservice will start running on port 8081 by default.

### Usage

To validate a credit card number, send a POST request to the `/validate` endpoint with a JSON payload containing the credit card number:

```bash
curl -X POST http://localhost:8081/validate -d '{"card_number":"1234567890123456"}' -H "Content-Type: application/json"
```

Replace `"1234567890123456"` with the credit card number you want to validate.

### API Response

The API response will be in JSON format:

```json
{
  "is_valid": true,
  "message": "Validation complete"
}
```