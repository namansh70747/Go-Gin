# 📝 Go Gin Notes API

A modern, production-ready RESTful API for managing notes, built with [Go](https://golang.org/) and [Gin](https://gin-gonic.com/en/docs/introduction/). This project demonstrates best practices for structuring Go web applications, using PostgreSQL with GORM, and following professional development workflows.

---

## 🚀 Features
- Create, read, and fetch notes (CRUD)
- PostgreSQL database integration via GORM
- Environment variable configuration with `.env`
- Production-ready Gin setup (release mode, trusted proxies, custom port)
- Clean, modular project structure
- Beginner-friendly code and documentation

---

## 📁 Project Structure
```
cmd/         # Entry point (main.go)
config/      # Database and app configuration
controllers/ # Request handlers (business logic)
models/      # Data models (structs for DB)
routes/      # Route definitions
middleware/  # (For custom middleware, e.g. auth)
services/    # (For reusable business logic)
repository/  # (For DB access logic)
utils/       # (For helper functions)
static/      # (For static files)
templates/   # (For HTML templates, optional)
.env         # Environment variables (not committed)
README.md    # Project documentation
```

---

## ⚡️ Quickstart

### 1. Clone the repo
```sh
git clone <your-repo-url>
cd Go-Gin
```

### 2. Install dependencies
```sh
go mod tidy
```

### 3. Set up PostgreSQL
- Start PostgreSQL: `brew services start postgresql`
- Create a database:
  ```sh
  psql postgres
  CREATE DATABASE notesdb;
  \q
  ```

### 4. Create a `.env` file in the project root:
```
GIN_MODE=release
PORT=8080
TRUSTED_PROXIES=127.0.0.1
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=notesdb
DB_PORT=5432
```

### 5. Run the app
```sh
go run cmd/main.go
```

---

## 📚 API Endpoints

| Method | Endpoint         | Description              |
|--------|------------------|--------------------------|
| GET    | /notes           | List all notes           |
| GET    | /notes/:id       | Get a note by ID         |
| POST   | /notes           | Create a new note        |

### Example POST Body
```json
{
  "title": "My First Note",
  "content": "This is the content of my first note."
}
```

---

## 🛠️ Environment Variables
| Variable         | Description                        |
|------------------|------------------------------------|
| GIN_MODE         | Gin mode (release/debug)           |
| PORT             | Port to run the server on          |
| TRUSTED_PROXIES  | Comma-separated list of proxies    |
| DB_HOST          | Database host                      |
| DB_USER          | Database user                      |
| DB_PASSWORD      | Database password                  |
| DB_NAME          | Database name                      |
| DB_PORT          | Database port                      |

---

## 🧑‍💻 Contributing
1. Fork this repo
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

---

## 🙏 Acknowledgements
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM ORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [godotenv](https://github.com/joho/godotenv)

---

## 💡 Tips
- For production, set `GIN_MODE=release` and configure trusted proxies.
- Use tools like [Postman](https://www.postman.com/) or `curl` to test your API.
- Keep your `.env` file secret!

---

Happy coding! 🚀
