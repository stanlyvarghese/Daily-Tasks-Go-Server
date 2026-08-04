# Go Multi-Page Web Server

A beginner-friendly web server built using Go's standard `net/http` package.

This project demonstrates how to create a simple HTTP server, register multiple routes, and serve HTML pages directly from Go without using any external frameworks.

---

## Features

- Built using only Go's standard library
- Multiple web pages
- HTTP routing
- HTML served directly from Go
- Navigation between pages
- Beginner-friendly code structure

---

## Pages

| Route | Description |
|--------|-------------|
| `/` | Home page |
| `/about` | About page |
| `/contact` | Contact page |

---

## Project Structure

```
go-multi-page-web-server
│
├── main.go
├── go.mod
├── README.md
├── LICENSE
└── .gitignore
```

---

## Concepts Practiced

This project helped me learn:

- Packages
- Functions
- HTTP Servers
- HTTP Handlers
- ResponseWriter
- Request
- Routing
- HTML in Go
- fmt package
- net/http package

---

## Technologies Used

- Go
- net/http
- HTML

No external packages or frameworks were used.

---

## Getting Started

### Clone the repository

```bash
git clone https://github.com/yourusername/go-multi-page-web-server.git
```

### Navigate into the project

```bash
cd go-multi-page-web-server
```

### Initialize Go modules (if needed)

```bash
go mod tidy
```

### Run the application

```bash
go run main.go
```

The server starts on:

```
http://localhost:8080
```

---

## Screens

### Home

- Navigation bar
- Task list
- Links to About and Contact pages

### About

Simple About page with navigation buttons.

### Contact

Simple Contact page displaying contact information.

---

## Learning Objectives

The goal of this project was to understand:

- How Go handles HTTP requests
- Registering routes using `http.HandleFunc`
- Writing responses using `fmt.Fprintln`
- Passing `http.ResponseWriter`
- Receiving requests with `*http.Request`
- Serving HTML directly from Go code
- Running a local web server

---

## Future Improvements

- Separate HTML into templates
- Use Go HTML templates
- Add CSS files
- Serve static assets
- Add forms
- Handle POST requests
- Use Templ
- Add HTMX
- Organize handlers into packages
- Add logging
- Improve UI

---

## Sample Output

```
Home
--------------------------------
Learn Go + Templ + HTMX

About

Contact
```

---

## What I Learned

While building this project, I learned:

- Creating a web server in Go
- Using the `net/http` package
- Registering multiple routes
- Returning HTML responses
- Handling browser requests
- Building a basic multi-page website without external libraries

---

## Author

**Stanly Varghese**

Aspiring Software Development Engineer

Learning Go one project at a time.

---

## Acknowledgements

This project was built as part of my Go learning journey to understand backend web development using Go's standard library.
