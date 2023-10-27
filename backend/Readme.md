<h1 align="center">CloudShareX - Backend</h1>

- [Technologies Utilized](#technologies-utilized)
- [Services](#services)
    - [Auth Service](#auth-service)
    - [File Service](#file-service)
    - [Email Service](#email-service)
    - [Storage Service](#storage-service)
- [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Configuration](#configuration)
- [Usage](#usage)
- [API Documentation](#api-documentation)

## Technologies Utilized

The project leverages a versatile stack of technologies, enhancing its capabilities and functionality:

- **Go**: The primary programming language that powers the project's backend services.
- **Firebase**: Used for cloud services and authentication.
- **MongoDB**: A NoSQL database for flexible data storage.
- **Postgres**: A robust relational database for structured data.
- **gRPC**: A high-performance, language-agnostic remote procedure call framework.
- **REST**: Representational State Transfer, an architectural style for web services.

These technologies, combined, empower the project to deliver a robust and scalable solution.


## Services

### Auth Service

The Auth Service is responsible for managing user authentication. It provides the following functionality:
- Creation of user accounts.
- User sign-in and authentication.
- Retrieving user information.

### File Service

The File Service is responsible for managing files, including uploading, retrieving, and tracking file history. It offers the following features:
- File upload and storage.
- Retrieval of uploaded files.
- File upload history tracking.
- Creates unique identifier with Google UUID for each file.
- Creates password and hash them with BCrypt to save safely.
### Email Service

The Email Service handles sending email notifications for various events, including:
- Account creation notifications.
- File expiration alerts.
- And more.

### Storage Service

The Storage Service periodically checks Firebase Storage to manage and clean up expired files. It performs the following tasks:
- Scans Firebase Storage for expired files.
- Deletes expired files to free up storage space.
