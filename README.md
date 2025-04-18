# 👤 Go Person CLI

A terminal-based contact manager built with Go.  
This application allows users to manage their contacts from the command line with persistent local storage using JSON.

---

## 🚀 Features

- Add, list, search, update, delete contacts
- `people.json` is automatically created on first save
- A backup file `backup_people.json` is generated every time you save
- Fallback mechanism loads from backup if the main file is corrupted
- Error handling with wrapped context using `fmt.Errorf`
- Interface-driven architecture with testable components
- `MemoryStore` used for unit testing (mock store)

---

## 📦 Installation

Make sure Go is installed on your machine.

```bash
git clone https://github.com/YOUR_USERNAME/go-person-cli.git
cd go-person-cli
go mod tidy
go run main.go
