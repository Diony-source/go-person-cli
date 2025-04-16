# 👤 Go Person CLI

A terminal-based contact manager built with Go.  
This application allows users to manage their contacts from the command line with persistent local storage using JSON.

---

## 🚀 Features

- Add, list, search, update, delete contacts
- `people.json` will be created automatically on first save
- A backup file `backup_people.json` will be generated each time you save
- Struct-based design for clean data representation
- Interface-driven architecture with testable components
- `MemoryStore` used for unit testing (mocked store)

---

## 📦 Installation

Make sure Go is installed on your machine.

```bash
git clone https://github.com/YOUR_USERNAME/go-person-cli.git
cd go-person-cli
go mod tidy
go run main.go
