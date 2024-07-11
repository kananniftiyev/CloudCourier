# CloudCourier 📦 - Backend

---

Welcome to the CloudCourier Backend, where the magic happens! 🧙‍♂️

- [CloudCourier 📦 - Backend](#cloudcourier----backend)
  - [Technologies Utilized](#technologies-utilized)
  - [Services](#services)
    - [🔒 Auth Service](#-auth-service)
    - [📂 File Service](#-file-service)
    - [📧 Email Service](#-email-service)
    - [☁️ Storage Service](#️-storage-service)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Configuration](#configuration)
    - [1. Firebase Credentials](#1-firebase-credentials)
    - [2. Database Settings](#2-database-settings)
  - [API Documentation](#api-documentation)

## Technologies Utilized

The project leverages a versatile stack of technologies, enhancing its capabilities and functionality:

- 🐍 **Go**: The primary programming language that powers the project's backend services.
- 🌐 **Firebase**: Used for cloud services and authentication.
- 🍃 **MongoDB**: A NoSQL database for flexible data storage.
- 🐘 **Postgres**: A robust relational database for structured data.
- 🌐 **REST**: Representational State Transfer, an architectural style for web services.

These technologies, combined, empower the project to deliver a robust and scalable solution. 🛠

## Services

### 🔒 Auth Service

The Auth Service is responsible for managing user authentication. It provides the following functionality:

- Creation of user accounts.
- User sign-in and authentication.
- Retrieving user information.

### 📂 File Service

The File Service is responsible for managing files, including uploading, retrieving, and tracking file history. It offers the following features:

- File upload and storage.
- Retrieval of uploaded files.
- File upload history tracking.
- Creates a unique identifier with Google UUID for each file.
- Creates passwords and securely hashes them with BCrypt for safe storage. 🔐

### 📧 Email Service

The Email Service handles sending email notifications for various events, including:

- Account creation notifications.
- File expiration alerts.
- And more.

### ☁️ Storage Service

The Storage Service periodically checks Firebase Storage to manage and clean up expired files. It performs the following tasks:

- Scans Firebase Storage for expired files.
- Deletes expired files to free up storage space. 🗑️

## Getting Started

### Installation

- Clone this Repository

```bash
git clone https://github.com/kananniftiyev/CloudCourier
```

- Move into service folder which you wanna run.

```bash
cd SERVICE_DIR
```

- Run **main.go** file in **cmd** folder to start service.

```bash
go run main.go
```

### Configuration

Before you begin working on this project, it's important to perform some initial configuration to ensure a smooth setup.

### 1. Firebase Credentials

**Location:** `backend` folder

1. Place your Firebase credentials file in the `` folder.
2. Update the file path to your Firebase credentials in the `.env` file to ensure proper authentication and access to Firebase services.

### 2. Database Settings

In each service's respective folder, make sure to configure the following settings in .env file:

- **Port**: Set the correct port for the database.
- **User**: Provide the username required for accessing the database.
- **Password**: Specify the password necessary for database access.

By adjusting these configuration parameters in the appropriate database folders for each service, you'll establish the correct database connectivity.

By following these configuration guidelines, you'll be well-prepared to initiate and run your project with all the necessary settings in place.
You can use this Markdown code to replace the content of your readme.md file.

## API Documentation

Api Documentation is in the folder of each Service that uses REST API. These Are:

- Auth Service
- File Service
