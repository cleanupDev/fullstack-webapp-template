# Full-Stack Web Application Template

---

## Table of Contents

- Overview
- Features
- Technologies Used
- Setup and Installation
  - Manual
  - Docker
- Usage
- API Documentation
- License
- Acknowledgments

## Overview

The Full-Stack Web Application Template is a learning project that demonstrates the implementation of a full-stack web application. The project consists of a backend written in Go (Golang) and a frontend built with Flask. The web application provides basic user authentication features, allowing users to register, login, logout, and access protected routes.

## Features

- User registration and login functionality.
- Secure password hashing using bcrypt for user authentication.
- Protected routes that require users to be logged in to access.
- Basic session management using Flask-Login for frontend authentication.

## Technologies Used

- Backend (Go):
  - Go (Golang) - The programming language for the backend.
  - Gin - A lightweight web framework for Go.
  - Go-MySQL-Driver - A MySQL-Driver for Go's database/sql package
  - bcrypt - A library for secure password hashing.

- Frontend (Flask):
  - Python - The programming language for the frontend.
  - Flask - A micro web framework for Python.
  - Flask-Login - A user session management for Flask.

- Docker
  - docker compose - For simple deployment and development.
  - mysql - Official Docker image for MySQL as a Database.
  - Adminer - For easy database management.

## Setup and Installation

### Manual

1. Clone the repository:

``` bash
    git clone https://github.com/cleanupDev/verbose-pancake.git
    cd verbose-pancake
```

2. Environment
   - Be sure to set up the .env files in the environments directory!

3. Backend Setup:

   - Install Go and set up your Go workspace.
   - Install the required Go dependencies:

     ``` bash
     go get -u github.com/gin-gonic/gin
     go get -u github.com/go-sql-driver/mysql
     go get -u github.com/joho/godotenv
     go get -u golang.org/x/crypto/bcrypt
     ```

   - Run the backend server:
     - main.go is located in `backend/cmd/main.go`

     ``` bash
     go run main.go
     ```

4. Frontend Setup:

   - Install Python and set up a virtual environment (optional but recommended).
   - Install the required Python dependencies:

     ``` bash
     pip install -r requirements.txt
     ```

   - Run the frontend server:

     ``` bash
     python app.py
     ```

5. Access the web application at `http://localhost:5000` in your web browser.

### Docker

 1. Environment
    - Set up the .env files in the environments directory!
 2. Docker compose
     - Simply run `docker compose up -d` to start the webapp.
 3. Access the web application at `http://localhost:5000` in your web browser.


## Usage

1. Register as a new user using the "Register" link on the homepage.
2. Log in with your registered credentials using the "Login" link.
3. After successful login, you can access the protected "Home" route.
4. Click on "Logout" to log out of the application.

## API Documentation

The backend provides a RESTful API for user registration and authentication. The API endpoints are as follows:

- `POST /create/user` : Register a new user. Accepts `username`, `email`, `password`, `first_name` and `last_name` in the request body.

- `POST /login`: Authenticate and log in a user. Accepts `username` and `password` in the request body.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

- [Go Programming Language](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)
- [Flask Microframework](https://flask.palletsprojects.com/)
- [Flask Login](https://github.com/maxcountryman/flask-login)

---

### WIP
