# 🚀 Nexcraft GoCore – REST API Starter Kit

**Nexcraft GoCore** is a clean, modern, and extensible REST API boilerplate built with **Golang**, **Fiber**, and **PostgreSQL** using `sqlx`. It’s designed to help developers start backend projects faster with a modular architecture and production-ready setup.

---

## ✨ Features

- ⚡ **Fiber Framework** – Fast, minimalist web framework for Go
- 🧱 **Modular Architecture** – Clean separation of concerns (Controllers, Services, Models)
- 🔐 **Environment Config** – Seamless config management with `.env`
- 🗃️ **PostgreSQL + sqlx** – Flexible and performant SQL layer
- 🧩 **Middleware-Ready** – Extend with logging, auth, rate limiting, etc.
- 🔁 **Starter CRUD** – User module (Create & Read included)

---

## 🛠️ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/nexcraft-gocore-starter-kit.git
cd nexcraft-gocore-starter-kit
```

### 2. Setup PostgreSQL

Create a database:
```bash
CREATE DATABASE nexcraft_dev;
```
Run the schema:
```bash
psql -U your_user -d nexcraft_dev -f scripts/init.sql
```

### 3. Configure Environment Variables

Copy the example config:
```bash
cp .env.example .env
```
Edit .env:
```bash
APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=nexcraft_dev
```
### 4. Run the Application

go run cmd/main.go

🔀 Sample API Routes
Method	Endpoint	Description
GET	    /api/users	Get all users
POST	/api/users	Create a new user

---

📦 Dependencies

    Fiber – Web framework

    sqlx – SQL toolkit for Go

    godotenv – .env file loader

---

📜 License

MIT License. You are free to use this for personal, educational, or commercial purposes.

---

💡 About Nexcraft

Nexcraft Studio builds practical developer tools and digital products to help startups and teams launch faster.