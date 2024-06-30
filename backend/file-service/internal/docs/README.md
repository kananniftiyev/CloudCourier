# 🚀 Golang REST Microservice - File Service

Welcome to the Golang REST Microservice for File Service! 📂 This powerful service allows you to manage and interact with files efficiently.

## 📁 Project Structure

The project is structured as follows:
- **Database**: Handles database-related functionality.
- **Rest**: Contains the REST API handlers and routes.
- **Tests**: Reserved for testing files, but you can extend it as needed.

## 🌐 REST API Endpoints

### File Upload 📤
- **POST /file/upload**
    - 📝 Description: Upload a file to the server.
    - 📦 Parameters:
        - `file` (multipart/form-data): The file to upload.
        - `password` (string): An optional password to protect the file.
        - `title` (string): An optional title for the file.
    - 📬 Returns: JSON response.

### File Retrieve 📥
- **GET /file/retrieve**
    - 📝 Description: Retrieve file information by its UUID.
    - 📦 Parameters:
        - `uuid` (string): The UUID of the file.
    - 📬 Returns: JSON response.

### File Upload History 🕒
- **GET /file/history**
    - 📝 Description: Retrieve a user's file upload history.
    - 📦 Parameters:
        - `username` (string): The username of the user.
    - 📬 Returns: JSON response.

## 🧩 Code Structure

Here's an overview of the code files within this microservice:

- **rest/file_handler.go**: Contains the HTTP request handlers for file upload, retrieval, and history. It also manages file storage in Firebase and MongoDB.
- **rest/file_models.go**: Defines the data models used by the REST handlers.
- **rest/file_rest.go**: Sets up the REST API server and routes using the Chi router.
- **rest/file_routes.go**: Initializes the routes for the REST API.
- **file/file.go**: Includes utility functions for interacting with Firebase and handling JWT tokens.
- **database/database.go**: Manages the MongoDB connection and environment variables.
- **database/models.go**: Defines the data structure for files.
- **database/file_repository.go**: Handles database operations for files.

## 🚧 Important Notes

- **TODO**: Some parts of the code are marked with TODO comments, indicating areas for improvement or future work.
- **Todo**: Specific tasks to be completed are marked with "Todo" comments.

## 🚀 Getting Started

To start the File Service, use the `FileStart` function in `rest/file_rest.go`. By default, the service listens on port `8081`. Ensure you configure Firebase and MongoDB connections before running the service.

You're all set to use this Golang REST Microservice for File Service! Happy coding! 🚀

---

For any questions or further assistance, feel free to ask. Enjoy building amazing applications with this microservice! 🌟
