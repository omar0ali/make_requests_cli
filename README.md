# HTTP Request Helper

## Description

`make_requests_cli` is a command-line interface (CLI) tool built in Go (Golang) that allows users to interactively create and manage HTTP requests. It provides a simple and intuitive way to perform CRUD (Create, Read, Update, Delete) operations on requests, as well as sending HTTP requests to a specified endpoint.

## Features

-   Create, read, and delete templates
-   Create, read, and delete requests
-   Send HTTP POST, GET, and DELETE requests
-   Interactive user interface for easy operation

## Setup

To use `make_requests_cli`, follow these steps:

1. Clone the repository or download the source code.

```bash
git clone https://github.com/omar0ali/make_requests_cli
```

2. Install Go if you haven't already (https://golang.org/doc/install).
3. Navigate to the project directory in your terminal.
4. Build the application by running `go build .`.
5. Run the application. `go run .`
6. Follow the prompts to interactively create and manage templates and requests, as well as send HTTP requests.

## Example Usage: Without Using Tview UI Design Improvement

Here's an example of how to use `make_requests_cli`:

1. Run the application.
2. Follow the prompts to create a template with a specified URL, port number, and HTTPS option.
3. Create a request using the created template and specify the path.
4. Select the request type (POST, GET, DELETE) and enter any required data.
5. Send the HTTP request and view the response.

## Example Usage: With Tview UI Design Improvement

(Under Development)

TView Library (https://github.com/rivo/tview)

To try a demo make the following changes to the `app.go` file.

```go
func main() {
	database.Connect()
	sqlDB, err := database.DB.DB()
	if err != nil {
		panic("Problem: with getting db.")
	}
	defer sqlDB.Close()

	RunApp(NEW)
    //RunApp(OLD)
}
```

### Whats not working:

-   HTTP GET
-   HTTP DELETE
-   HTTP UPDATE

## Notes

-   The application provides guidance and reminders throughout the process, such as reminding users that field names are case-sensitive and should match the backend.
-   There are multiple ways to close the application: you can choose the '0' option to exit, or you can close it by pressing CTRL+C.
-   **Note**: This app might not be easy or convenient to use for everyone, although I created it to make my workflow easier for testing my backend API and other tasks.

## Dependencies

-   Golang (https://golang.org/)
-   External dependencies managed via Go modules

> [!NOTE]
> This project is developed as a learning exercise in Golang. While efforts have been made to ensure its functionality, there may be bugs or errors present.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
