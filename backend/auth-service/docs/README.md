# 🚀 Golang REST Microservice - Auth Service

Welcome to the Golang REST Microservice for Auth Service! 🛡️ This service enables user registration, login, and authentication.

## 📁 Project Structure

The project is organized into the following directories:
- **Database**: Manages the database operations for user data.
- **Rest**: Contains the REST API handlers and routes for authentication.
- **Docs**: Documentation (not included in this Markdown file).
- **Tests**: Reserved for testing, but you can extend it as needed.

## 🗃 Database Models

### User Model
- Stores user information, including username, email, password hash, and registration date.

## 🛠 Database Repository

The **UserRepository** handles database operations, including user creation, login checks, and retrieving user data.

## ⚙ Database Configuration

The **ConnectDatabase** function establishes the database connection using environment variables.

## 🌐 REST API Endpoints

### Register User 📝
- **POST /auth/register**
    - 📝 Description: Register a new user.
    - 📦 Parameters: JSON request with username, email, and password.
    - 📬 Returns: JSON response.

### Login User 🔐
- **POST /auth/login**
    - 📝 Description: Authenticate and log in a user.
    - 📦 Parameters: JSON request with email and password.
    - 📬 Returns: JSON response and a JWT token.

### Logout User 🚪
- **POST /auth/logout**
    - 📝 Description: Log out the authenticated user.
    - 📬 Returns: JSON response.

### Get User Data 📧
- **GET /auth/user**
    - 📝 Description: Retrieve user data.
    - 📦 Requires: JWT token for authentication.
    - 📬 Returns: JSON response with user information.

## 🧩 Code Structure

Here's an overview of the code files within this microservice:

- **database/models/models.go**: Defines the data model for user entities.
- **database/repository/user_repository.go**: Manages database operations for users and handles user-related errors.
- **database/database.go**: Establishes the database connection and initializes the database.
- **rest/auth_handlers.go**: Contains the HTTP request handlers for user registration, login, logout, and user data retrieval.
- **rest/auth_models.go**: Defines the data structures for authentication-related requests and responses.
- **rest/auth_rest.go**: Sets up the REST API server and routes using the Chi router.
- **rest/routes.go**: Initializes the routes for the REST API.
- **auth/auth.go**: Provides authentication-related functions such as hashing passwords, verifying passwords, creating JWT tokens, and handling JWT tokens.

## 🚀 Getting Started

To start the Auth Service, use the `AuthStart` function in `rest/auth_rest.go`. By default, the service listens on port `8080`. Ensure you configure the database connection and secret key before running the service.

You're ready to use this Golang REST Microservice for Auth Service! Safeguard your applications with user authentication and authorization. 🚀

---

For any questions or further assistance, feel free to ask. Enjoy building secure applications with this microservice! 🌟
