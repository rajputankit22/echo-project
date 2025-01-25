# Echo Project

This project is built using the Echo framework for Go.

## Prerequisites

- Go 1.16 or higher
- Echo v4

## Installation

1. Clone the repository:
    ```sh
    git clone git@github.com:rajputankit22/echo-project.git
    ```
2. Navigate to the project directory:
    ```sh
    cd echo-project
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Running the Project

To run the project, use the following command:
```sh
go run main.go
```

## Environment Variables

The following environment variables are used in this project:
you can copy example.env and past in your project's .env file.

- `HTTP_PORT`: The port on which the server will run (default: `8080`).
- `DB_HOST`: The hostname of the database server.
- `DB_USER`: The username for the database.
- `DB_PASSWORD`: The password for the database.
- `DB_NAME`: The name of the database.
- `LOG_LEVEL`: The name of the database.

## Project Structure

```
/echo-project
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.