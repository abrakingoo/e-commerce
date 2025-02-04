# E-Commerce Web Application

## Project Overview
A full-featured e-commerce web application built with Go, featuring user authentication, product browsing, cart management, and admin functionalities.

#### live link : https://skymall.onrender.com

## Features
- User Authentication (Sign Up, Sign In, Sign Out)
- Product Browsing
- Shopping Cart Management
- Checkout Process
- Admin Dashboard
- Hot Deals Section
- Product Upload and Stock Management
- User Profile Management

## Technology Stack
- Language: Go (Golang)
- Database: SQLite
- Web Server: Standard Go `net/http` package
- Authentication: Custom session-based authentication

## Project Structure
```
e-commerce/
│
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency lock file
│
├── db/                     # Database initialization and migrations
├── handlers/               # HTTP route handlers
│   ├── admin.go
│   ├── cart.go
│   ├── checkout.go
│   ├── product.go
│   └── ...
│
├── templates/              # HTML templates
├── static/                 # Static assets (CSS, JS, images)
├── utils/                  # Utility functions
└── data/                   # Data-related operations
```

## Prerequisites
- Go 1.16+ installed
- SQLite3

## Setup and Installation
1. Clone the repository
   ```bash
   git clone https://github.com/abrakingoo/e-commerce/
   cd e-commerce
   ```

2. Install dependencies
   ```bash
   go mod tidy
   ```

3. Run the application
   ```bash
   go run main.go
   ```

4. Access the application
   - Open a web browser
   - Navigate to `http://localhost:10000`

## Available Routes
- `/`: Home page
- `/signin`: User login
- `/signup`: User registration
- `/signout`: User logout
- `/cart`: Shopping cart
- `/checkout`: Checkout process
- `/profile`: User profile
- `/product`: Product listing
- `/hotdeals`: Special deals
- `/admin`: Admin dashboard
- `/upload`: Product upload (admin)
- `/stock`: Stock management (admin)

Admin Credentials
   ```bash
   email:  admin@skymal.com
   password: 1234
   ```

## Database
- Uses SQLite for data storage
- Database file: `database.db`
- Automatically creates necessary tables on first run

## Security Features
- Session-based authentication
- Protected admin routes
- Input validation
- Static file serving with directory browsing prevention

## Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## Contact
maikoabraham68@gmail.com