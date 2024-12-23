# Registration Service API

This project provides an API for registering users with a PAN number, mobile number, email, and name. It uses the Gin framework, Go's standard library, and Docker for easy containerization.

## Prerequisites

- Docker installed on your system.
- Go 1.22.3+ installed on your local machine (if you want to run the code locally).

## Getting Started

1. **Clone the repository**:

    ```bash
    git clone https://github.com/rishabh625/registration-service.git
    cd registration-service
    ```

2. **Running with Docker**:

    - Build and run the Docker container:

        ```bash
        docker build -t registration-service .
        docker run -p 8080:8080 registration-service
        ```

    - The server will be available on `http://localhost:8080`.

3. **Running locally** (without Docker):

    - Install dependencies:

        ```bash
        go mod tidy
        ```

    - Run the application:

        ```bash
        go run server/cmd/main.go 
        ```

    - The server will be available on `http://localhost:8080`.

## API Endpoints

### `POST /register`
Registers a new user with the provided PAN, mobile number, email, and name.

#### Request Body:

```json
{
    "pan": "ABCDE1234F",
    "number": "9876543210",
    "name": "John Doe",
    "email": "john@example.com"
}
```
## Test API

Register a new User
Curl:

```
curl --location 'http://localhost:8080/users/register' \
--header 'Content-Type: application/json' \
--data-raw '{
"name": "Rishabh",
"pan": "AUKLE2889K",
"number": "9890999990",
"email": "rishabh@gmail.com"
}'
```

Verify User
Curl:

```
curl --location --request GET 'http://localhost:8080/users/CUKPM2889K' \
--header 'Content-Type: application/json' \
--data-raw '{
"name": "Rishabh Mishra",
"pan": "AUKLE2889K",
"number": "9890999990",
"email": "rishabh@gmail.com"
}'
```
