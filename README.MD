
# Go Bookstore API
This is a simple RESTful API for managing a bookstore, built with Go. It allows you to perform CRUD operations on books in a SQLite database.

## Project Structure
go-bookstore/
├── pkg/
│   ├── config/
│   │   └── app.go
│   ├── controllers/
│   │   ├── book_controller.go
│   │   └── book_controller_test.go
│   ├── models/
│   │   ├── book.go
│   │   └── book_test.go
│   ├── utils/
│   │   └── utils.go
├── main.go
├── go.mod
├── go.sum
└── README.md


## Requirements

- Go 1.16 or later
- Postgres 

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/go-bookstore.git
cd go-bookstore
```

2. Install dependencies:

```sh
go mod tidy
```

## Running the API

To run the API server:

```sh
go run .\cmd\main\main.go
```

The server will start on `http://localhost:9010`.

## API Endpoints

- `POST /book`: Create a new book
- `GET /books`: Get all books
- `GET /book/{bookID}`: Get a book by its ID
- `PUT /book/{bookID}`: Update a book by its ID
- `DELETE /book/{bookID}`: Delete a book by its ID

## Project Files

### `app.go`

Contains the configuration for connecting to the database.

### `book_controller.go`

Contains the HTTP handlers for the book-related endpoints.

### `book.go`

Contains the model definition and methods for interacting with the `Book` table in the database.

### `utils.go`

Contains utility functions, such as `ParseBody` for parsing request bodies.

## Example Requests

Here are some example requests you can use to interact with the API:

### Create a Book

```sh
curl -X POST http://localhost:9010/book -d '{
    "name": "Go Programming",
    "author": "John Doe",
    "publication": "Tech Books"
}' -H "Content-Type: application/json"
```

### Get All Books

```sh
curl http://localhost:9010/books
```

### Get a Book by ID

```sh
curl http://localhost:9010/book/1
```

### Update a Book

```sh
curl -X PUT http://localhost:9010/book/1 -d '{
    "name": "Go Programming Updated",
    "author": "Jane Doe",
    "publication": "Tech Books Updated"
}' -H "Content-Type: application/json"
```

### Delete a Book

```sh
curl -X DELETE http://localhost:9010/book/1
```

## Contributing

Contributions are welcome! Please create an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
