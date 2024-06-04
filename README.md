# Go Simple Web Server

This is a basic web server written in Go that demonstrates handling static files, processing GET requests, and handling form submissions.

## Description

This server serves static files from the `./static` directory. It also provides two functionalities:

1. **`/`**: Responds with "Homepage".
2. **`/hello`**: Responds with "hello!" for GET requests to this path.
3. **`/form.html`**: Handles form submissions. It parses the submitted data and echoes the values of "name" and "address" back to the user.

## Technology Used

- Golang

## Getting Started

### Prerequisites

Make sure you have Golang installed on your machine.

### Installation

1. Clone the repository:

```bash
git clone https://github.com/Rupesh180902/basic-server-with-golang.git
```

2. Navigate to the project directory:

```bash
cd server
```

3. build the project:

```bash
go build
```

**Usage**

1. Start the development server:

```bash
go run main.go
```

2. Open your browser and visit http://localhost:8080 .

```bash
go run server.go
```

This will start the server on port 8080.

### There are 3 endpoints of the server:

- http://localhost:8080
- http://localhost:8080/hello
- http://localhost:8080/form.html
