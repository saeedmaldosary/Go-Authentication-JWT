# Go-Authentication-JWT

A simple Go application that demonstrates user authentication using JWT (JSON Web Tokens).

## Features

- **User Signup**: Allows new users to register.
- **User Login**: Authenticates users and provides a JWT token.
- **Token Validation**: Validates the provided JWT token.

## Setup and Initialization

1. **Environment Variables**: Ensure you have a `.env` file with the necessary environment variables. The application uses `godotenv` to load these variables.
   
2. **Database Connection**: The application connects to a PostgreSQL database. The connection string is fetched from the environment variable `DB`.

3. **Database Synchronization**: The application uses GORM for ORM and automatically migrates the user model to the database.

## Endpoints

- **Signup**: `POST /signup`
  - Payload: 
    ```json
    {
      "Email": "example@email.com",
      "Password": "your_password"
    }
    ```
  
- **Login**: `POST /login`
  - Payload:
    ```json
    {
      "Email": "example@email.com",
      "Password": "your_password"
    }
    ```
  - Response: Returns a JWT token.

- **Validate**: `GET /validate`
  - Requires JWT token in the `Authorization` cookie. Validates the token and returns the associated user.

## Middleware

- **RequireAuth**: This middleware checks for the `Authorization` cookie in the request, decodes the JWT token, validates it, and attaches the associated user to the request.

## Models

- **User Model**: Represents a user with fields `Email` and `Password`. The email is unique for each user.

