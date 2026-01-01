# 📚 Go Digital Library

A backend service for managing a digital library system, built with **Go**, **Gorilla Mux**, and **GORM**.  
It provides RESTful APIs for user authentication, book management, and loan tracking.


## 🚀 Features
- JWT‑based authentication middleware
- Role‑based access control (`admin` vs `user`)
- Borrow and return book flows with loan history
- Soft delete (archive) books using `is_active` flag
- Search books by title, author, and genre
- PostgreSQL integration with GORM ORM
- Service + Repository architecture for clean separation of concerns


## 🛠️ Tech Stack
- **Language:** Go 
- **Frameworks:** Gorilla Mux (routing), GORM (ORM)
- **Database:** PostgreSQL
- **Auth:** JWT (JSON Web Tokens)
- **Architecture:** Handlers → Services → Repository → Database


## 📂 Project Structure

/handlers      → HTTP handlers (API endpoints)  
/services      → Business logic (borrow/return flows)  
/repository    → Database queries with GORM  
/models        → Data models (Book, User, Loan)  
/middleware    → JWT authentication middleware  
/db            → Database connection setup  
main.go        → Application entrypoint   


## 🔑 API Endpoints

### Auth
- `POST /login` → Authenticate user, return JWT

### Books
- `GET /books` → List active books (filters: title, author, genre)
- `GET /books/user` → Get details of a user’s borrowed books
- `POST /books/{id}/borrow` → Borrow a book
- `POST /books/{id}/return` → Return a book
- `PATCH /books/{id}/archive` → Archive a book (admin only)

---

From Team Go Digital 

