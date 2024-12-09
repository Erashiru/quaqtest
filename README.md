# Weather API Service

This project is a test application. The goal is to build a REST API service using Go and MongoDB for managing weather data. The service will allow retrieving and updating weather information.

## Features

- **GET /weather**: Fetch weather data for a specific city from the database.
- **PUT /weather**: Update weather data in the database by fetching current weather information from an external API.

---

## Setup and Installation

1. **Clone the repository:**
   ```bash
   git clone <repository_url>
   cd <repository_directory>
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**
   Create a `.env` file in the root directory with the following content:
   ```plaintext
    MONGO_URI=
    DATABASE_NAME=
    WEATHER_API_KEY=
   ```

4. **Run the application:**
   ```bash
   go run ./cmd/
   ```

5. **Run tests:**
   ```bash
   go test ./...
   ```

---
