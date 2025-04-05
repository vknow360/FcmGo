# FcmGo

FcmGo is a lightweight Go application designed to send push notifications and messages using Firebase Cloud Messaging (FCM). It provides a simple REST API built with the Gin framework to send notifications to topics or specific device tokens.

## Features

-   Send notifications to FCM topics with customizable title, body, sound, and image.
-   Send data messages to FCM topics.
-   Send notifications to multiple device tokens.
-   Send data messages to multiple device tokens.
-   Simple status endpoint to check server health.
-   Environment variable configuration using `.env` files.

## Prerequisites

-   Go 1.21 or higher
-   Firebase project with a service account key (JSON file)
-   Git (optional, for cloning the repository)

## Installation

1. **Clone the Repository** (if applicable):

    ```bash
    git clone https://github.com/yourusername/fcmgo.git
    cd fcmgo
    ```

2. **Install Dependencies**:
   Run the following command to download all required Go modules:

    ```bash
    go mod tidy
    ```

3. **Set Up Environment Variables**:
   Create a `.env` file in the root directory and add your Firebase service account JSON:

    ```plaintext
    SERVICE_ACCOUNT_JSON={"type": "service_account", "project_id": "your-project-id", ...}
    ```

    Replace the value with your actual service account JSON string from the Firebase Console.

4. **Build and Run**:
   Compile and start the server:
    ```bash
    go build -o fcmgo
    ./fcmgo
    ```
    The server will run on `http://localhost:8080`.

## Usage

FcmGo exposes a REST API to interact with Firebase Cloud Messaging. Below are the available endpoints:

### Endpoints

#### 1. Check Server Status

-   **Endpoint**: `GET /status`
-   **Description**: Returns the server status.
-   **Response**:
    ```json
    {
        "status": "OK"
    }
    ```

#### 2. Send Notification to a Topic

-   **Endpoint**: `POST /send/topic/notification`
-   **Description**: Sends a notification to a specified FCM topic.
-   **Request Body**:
    ```json
    {
        "title": "Test Notification",
        "body": "This is a test message",
        "topic": "news",
        "sound": "default",
        "image": "https://example.com/image.png",
        "data": {
            "key1": "value1",
            "key2": "value2"
        }
    }
    ```
-   **Response**:
    ```json
    {
        "messageId": "projects/your-project-id/messages/123456789"
    }
    ```

#### 3. Send Message to a Topic

-   **Endpoint**: `POST /send/topic/message`
-   **Description**: Sends a data-only message to a specified FCM topic.
-   **Request Body**:
    ```json
    {
        "topic": "news",
        "data": {
            "key1": "value1",
            "key2": "value2"
        }
    }
    ```
-   **Response**:
    ```json
    {
        "messageId": "projects/your-project-id/messages/123456789"
    }
    ```

#### 4. Send Notification to Tokens

-   **Endpoint**: `POST /send/tokens/notification`
-   **Description**: Sends a notification to multiple device tokens.
-   **Request Body**:
    ```json
    {
        "title": "Test Notification",
        "body": "This is a test message",
        "tokens": ["token1", "token2"],
        "sound": "default",
        "image": "https://example.com/image.png",
        "data": {
            "key1": "value1",
            "key2": "value2"
        }
    }
    ```
-   **Response**:
    ```json
    {
        "successCount": 2,
        "failureCount": 0
    }
    ```

#### 5. Send Message to Tokens

-   **Endpoint**: `POST /send/tokens/message`
-   **Description**: Sends a data-only message to multiple device tokens.
-   **Request Body**:
    ```json
    {
        "tokens": ["token1", "token2"],
        "data": {
            "key1": "value1",
            "key2": "value2"
        }
    }
    ```
-   **Response**:
    ```json
    {
        "successCount": 2,
        "failureCount": 0
    }
    ```

## Project Structure

```
fcmgo/
├── config/          # Environment variable loading
├── firebase/        # Firebase initialization
├── handlers/        # API endpoint handlers
├── interfaces/      # Data types and structs
├── routes/          # API route definitions
├── utils/           # Logging utilities
├── main.go          # Application entry point
├── go.mod           # Go module dependencies
└── .env             # Environment variables (not tracked)
```

## Dependencies

-   [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
-   [Firebase Go SDK](https://firebase.google.com/docs/admin/setup#go) - Firebase integration
-   [Godotenv](https://github.com/joho/godotenv) - Environment variable loading

See `go.mod` for the full list of dependencies.

## Error Handling

-   If the `.env` file is missing or invalid, the application will log an error and exit.
-   API endpoints return appropriate HTTP status codes (e.g., `400` for bad requests, `500` for server errors) with error messages in the response body.
