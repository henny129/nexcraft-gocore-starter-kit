🚀 Nexcraft GoCore – REST API Starter Kit

Nexcraft GoCore is a modern, clean, and extensible REST API boilerplate built with Golang, Fiber, and PostgreSQL using sqlx. It helps you bootstrap your Go backend projects faster with production-ready architecture and essential features out of the box.
✨ Features

    ⚡ Fiber Framework – Blazing-fast HTTP web framework

    🧱 Modular Architecture – Clean and maintainable project structure

    🔐 Environment Config – Easy switching with .env support

    🗃️ PostgreSQL + sqlx – Structured SQL with flexibility

    🧩 Middleware-Ready – Easily add auth, logging, rate limiting, etc.

    🔁 Starter CRUD – User module with Create and Read already implemented

📁 Project Structure

nexcraft-gocore-starter-kit/
│
├── cmd/           # App entrypoint (main.go)
├── config/        # Env loader
├── controllers/   # HTTP handlers (for routing logic)
├── database/      # PostgreSQL connection setup
├── internal/
│   ├── models/    # Data models
│   └── services/  # Business logic layer
├── routes/        # API route definitions
├── scripts/       # SQL initialization scripts
├── utils/         # Helper functions (coming soon)
├── .env           # Environment configuration
├── go.mod         # Go module definition
└── go.sum         # Go module checksums

🛠️ Getting Started
1. Clone the Repository

git clone https://github.com/yourname/nexcraft-gocore-starter-kit.git
cd nexcraft-gocore-starter-kit

2. Setup PostgreSQL

Create a new database:

CREATE DATABASE nexcraft_dev;

Run the initialization script:

psql -U your_user -d nexcraft_dev -f scripts/init.sql

3. Configure Environment Variables

Copy .env.example to .env:

cp .env.example .env

Edit the .env file:

APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=nexcraft_dev

4. Run the App

go run cmd/main.go

📦 Dependencies

    Fiber – Web framework

    sqlx – Database layer

    godotenv – Loads .env files into os.Environ

📜 License

This project is licensed under the MIT License.
Feel free to use it for personal or commercial purposes.
💡 About Nexcraft

Nexcraft Studio builds practical developer tools and digital products to help startups, developers, and businesses launch faster.
