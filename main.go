package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

type PersonStore interface {
	Load() ([]Person, error)
	Save([]Person) error
}

type FileStore struct {
	FilePath string
}

func (f FileStore) Load() ([]Person, error) {
	file, err := os.ReadFile(f.FilePath)
	if err != nil {
		fmt.Println("📭 No data file found. Starting fresh.")
		return []Person{}, nil
	}

	var people []Person
	err = json.Unmarshal(file, &people)
	if err == nil {
		return people, nil
	}

	fmt.Println("⚠️ Main file corrupted. Trying backup...")

	backup, err := os.ReadFile("backup_people.json")
	if err != nil {
		fmt.Println("❌ Backup file missing.")
		return []Person{}, nil
	}

	err = json.Unmarshal(backup, &people)
	if err != nil {
		fmt.Println("❌ Backup file also broken. Starting empty.")
		return []Person{}, nil
	}

	fmt.Println("✅ Loaded from backup.")
	return people, nil
}

func (f FileStore) Save(people []Person) error {
	data, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(f.FilePath, data, 0644)
	if err != nil {
		return err
	}

	// 🔧 Backup file
	_ = os.WriteFile("backup_people.json", data, 0644)
	return nil
}

func main() {
	var store PersonStore = FileStore{FilePath: "people.json"}

	people, _ := store.Load()

	for {
		fmt.Print("\nEnter command (add, list, update, delete, search, save, exit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch strings.ToLower(cmd) {
		case "add":
			var p Person
			fmt.Print("Name: ")
			fmt.Scanln(&p.Name)
			fmt.Print("Age: ")
			fmt.Scanln(&p.Age)
			fmt.Print("Phone: ")
			fmt.Scanln(&p.Phone)
			people = append(people, p)
			fmt.Println("✅ Person added.")

		case "list":
			if len(people) == 0 {
				fmt.Println("📭 No people found.")
			}
			for i, p := range people {
				fmt.Printf("%d. %s (%d) - %s\n", i+1, p.Name, p.Age, p.Phone)
			}

		case "update":
			var index int
			fmt.Print("Enter index to update: ")
			fmt.Scanln(&index)
			if index < 1 || index > len(people) {
				fmt.Println("❌ Invalid index.")
				break
			}
			var updated Person
			fmt.Print("New Name: ")
			fmt.Scanln(&updated.Name)
			fmt.Print("New Age: ")
			fmt.Scanln(&updated.Age)
			fmt.Print("New Phone: ")
			fmt.Scanln(&updated.Phone)
			people[index-1] = updated
			fmt.Println("🔄 Updated.")

		case "delete":
			var index int
			fmt.Print("Enter index to delete: ")
			fmt.Scanln(&index)
			if index < 1 || index > len(people) {
				fmt.Println("❌ Invalid index.")
				break
			}
			people = append(people[:index-1], people[index:]...)
			fmt.Println("🗑️ Deleted.")

		case "search":
			var query string
			fmt.Print("Search by name: ")
			fmt.Scanln(&query)
			found := false
			for _, p := range people {
				if strings.EqualFold(p.Name, query) {
					fmt.Printf("🔍 %s (%d) - %s\n", p.Name, p.Age, p.Phone)
					found = true
				}
			}
			if !found {
				fmt.Println("🔎 Not found.")
			}

		case "save":
			err := store.Save(people)
			if err != nil {
				fmt.Println("❌ Save error:", err)
			} else {
				fmt.Println("💾 Saved.")
			}

		case "exit":
			fmt.Println("👋 Exit.")
			return

		default:
			fmt.Println("❓ Unknown command.")
		}
	}
}
