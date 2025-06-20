#  User Profile API Server

A full-stack web application built with **Go**, **PostgreSQL**, **React**, and **Tailwind CSS**. Users can be added via an API, upload their image, and generate a PDF profile. The frontend connects seamlessly with the backend and offers a sleek user interface.

---

## 🚀 Features

- ✅ Create, Read, Update, Delete (CRUD) users
- 📸 Upload user images
- 🗃 Store image and user data in PostgreSQL
- 📄 Generate user profile as a downloadable PDF
- 🎨 Modern, styled frontend built with React + Tailwind CSS
- 🔥 Backend and frontend fully connected

---

## 🧰 Tech Stack

**Frontend:**

- React.js  
- Tailwind CSS

**Backend:**

- Go (Golang)  
- PostgreSQL  
- net/http  
- goroutines for async PDF & image handling

---


---

## ⚙️ Getting Started

### 1. Start the Backend (Go)

```bash
cd go-user-service
go run main.go
```
### 2. Start the Frontend (React)

```bash
cd frontend
npm install
npm start
```
### API Endpoints

```bash
| Method | Endpoint       | Description          |
| ------ | -------------- | -------------------- |
| POST   | `/users`       | Add a new user       |
| GET    | `/users`       | Get all users        |
| GET    | `/users/:id`   | Get user by ID       |
| PUT    | `/users/:id`   | Update user name     |
| DELETE | `/users/:id`   | Delete user          |
| GET    | `/profile/:id` | Download PDF profile |
```

###  UI Preview
<img width="953" alt="image" src="https://github.com/user-attachments/assets/cb91e74b-5ab9-49b3-803c-d9456804bac2" />


<img width="959" alt="image" src="https://github.com/user-attachments/assets/02ddd644-f4db-4740-9594-a3db467ebe10" />

###  Author

Isha Raj 

GitHub: @isha54ry
Built during a custom API server project challenge 

### 📝 License
This project is for educational use. You can adapt or extend it as needed!


