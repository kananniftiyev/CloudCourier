# 🚀 Golang REST Microservice - File Service

Welcome to the Golang REST Microservice for File Service! 📁 This service enables user registration, login, and authentication.

## 📁 Project Structure

The project is organized into the following directories:

- **Database**: Manages the database operations for user data.
- **Rest**: Contains the REST API handlers and routes for authentication.
- **Tests**: Reserved for testing, but you can extend it as needed.

## 🗃 Database Models

### User Model

- Stores user information, including username, email, password hash, and registration date.

## 🛠 Database Repository

The **FileRepository** handles database operations, including user creation, login checks, and retrieving user data.

## ⚙ Database Configuration

The **ConnectDatabase** function establishes the database connection using environment variables.

## 🌐 REST API Endpoints

### Upload File 📤

- **POST api/file/upload**
  - 📝 Description: Upload new file.
  - 📦 Parameters: File and password as form value.
  - 📬 Returns: JSON response.

### Download File 📥

- **POST api/file/retrieve**
  - 📝 Description: Download file.
  - 📦 Parameters: UUID of File and password.
  - 📬 Returns: JSON response and a download link.

### File Upload History 📜

- **POST api/file/history**
  - 📝 Description: Review file upload history.
  - 📬 Returns: JSON response.

## 🚀 Getting Started

To start the File Service, use the `FileStart` function in `rest/rest.go`. By default, the service listens on port `8080`. Ensure you configure the database connection and secret key before running the service.

You're ready to use this Golang REST Microservice for File Service! Safeguard your applications with user authentication and authorization. 🚀

---

For any questions or further assistance, feel free to ask. Enjoy building secure applications with this microservice! 🌟
