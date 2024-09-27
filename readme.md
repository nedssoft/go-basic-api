# Go BASIC API

This project is a RESTful API built with Go, using Gin as the web framework and GORM for database operations.


## Features

- User management (create, read)
- Post management (create, read)
- Database migrations
- End-to-end testing

## Prerequisites

- Go 1.16+
- PostgreSQL

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/nedssoft/go-basic-api.git
   cd go-basic-api
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Set up your environment variables in a `.env` file:
   ```
   DB_HOST=localhost
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database_name
   DB_PORT=5432
   PORT=8080
   ```

4. Run database migrations:
   ```
   make migrate
   ```

## Running the Application

To start the server:
```
make run
```

The server will run on port 8080 by default.

## Testing

To run the tests:
```
make test
```

This will run the tests for both the user and post endpoints.


## API Endpoints

- `POST /api/v1/users`: Create a new user
- `GET /api/v1/users/:id`: Get a user by ID
- `POST /api/v1/posts`: Create a new post
- `GET /api/v1/posts`: Get all posts
- `GET /api/v1/posts/:id`: Get a post by ID

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
