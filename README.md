# LIBR_Prototype

A minimal version of LIBR system that simulates decentralized moderation, integrating basic APIs for message submission and retrieval.

---

## 📁Folder Structure

```
LIBR Prototype/
├── controller/
│   └── controller.go           # Handles HTTP requests
├── db/
│   └──db.go                    # DB connection
├── model/
│   └── model.go                # Data models
├── moderator/
│   └── moderator.go            # Moderation simulation
├── router/
│   └── router.go               # API Routes
├── .env                        # Environment variables
├── go.mod                      # Module file
├── main.go                     # Entry point
└── README.md
```

---

## 🧰Setup
Follow these steps to set up and run the project locally.

### 1️⃣Install Go
Download and install Go from: https://golang.org/dl/

Ensure Go is installed:

```bash
go version
```

### 2️⃣Install PostgreSQL

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

### 3️⃣Create database

```sql
CREATE DATABASE libr;
```

---

## 🧰Installation

### 1️⃣Clone repository

```bash
git clone <your-repo-url>
cd Libr_Prototype
go mod tidy
```

### 2️⃣Add environment file
Create a .env file in root directory

```env
DB_URL=postgres://<username>:<your_password>@localhost:5432/libr
```

Server will start on `http://localhost:8000`

---

## 📬API Endpoints

### 📨POST /post
Submit a message for moderation

**Payload**

```json
{ "content": "This is a prototype." }
```

**Response**

```json
{
  "id": "882c85b7-f92e45ca-8715-59eba12bc04d",
  "timestamp": 1749920048,
  "status": "Approved"
}
```

### 🔎GET /get/{timestamp}
Retrieve messages by timestamp

**Response**

```json
[
  {
    "id": "87adb610-47da-404b-9ab4-170302784ff2",
    "content": "Alphabets",
    "timestamp": 1749914634,
    "status": "Approved"
  }
]
```

### 📜GET /getall
Retrieve all messages stored in the database

---

## Screenshots
### ✅Approved Message
![ApprovedMessage](https://github.com/user-attachments/assets/551efb94-43d2-4f07-8904-e4527b61ede9)

### ❎Rejected Message
![RejectedMessage](https://github.com/user-attachments/assets/e33a8a98-da59-4b3d-9502-6fb416a3ed50)

### 🔍Get message with timestamp
![GetMessage](https://github.com/user-attachments/assets/fbf9345c-0cf7-4c1c-add2-f4de81459242)

### 📂Get all messages
![GetAllMessages](https://github.com/user-attachments/assets/f51ece33-fccd-40ed-928d-3313dfb098b5)

---
