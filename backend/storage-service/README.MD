# Storage Service

This is a Go application that performs storage operations using Firebase Storage. It includes a function to check and delete expired files from the storage bucket.

## Prerequisites

Before running this application, make sure you have the following:

- Go installed on your machine
- Firebase project set up with a storage bucket
- Environment variables configured (see `.env`)

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/storage-service.git
   ```

2. Navigate to the project directory:

   ```shell
   cd storage-service
   ```

3. Install the dependencies:

   ```shell
   go mod download
   ```

4. Update `.env` the values with your Firebase project details.

## Usage

To start the storage check and deletion process, run the following command:

```bash
go run main.go
```
