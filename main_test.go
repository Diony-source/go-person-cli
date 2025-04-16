package main

import (
	"testing"
)

type MemoryStore struct {
	Data []Person
}

func (m *MemoryStore) Save(people []Person) error {
	m.Data = people
	return nil
}

func (m *MemoryStore) Load() ([]Person, error) {
	return m.Data, nil
}

// 1. SAVE + LOAD test — correctness test
func TestMemoryStore_SaveAndLoad(t *testing.T) {
	store := &MemoryStore{}

	original := []Person{
		{Name: "X", Age: 99, Phone: "error"},
		{Name: "Y", Age: 0, Phone: "none"},
	}

	err := store.Save(original)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := store.Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if len(loaded) != len(original) {
		t.Errorf("Expected %d people, got %d", len(original), len(loaded))
	}

	for i := range loaded {
		if loaded[i] != original[i] {
			t.Errorf("Mismatch at index %d: got %+v, want %+v", i, loaded[i], original[i])
		}
	}
}

// 2. LOAD empty list test — edge-case test
func TestMemoryStore_LoadEmpty(t *testing.T) {
	store := &MemoryStore{}

	people, err := store.Load()
	if err != nil {
		t.Fatalf("Unexpected error during Load: %v", err)
	}

	if len(people) != 0 {
		t.Errorf("Expected empty slice, got %d people", len(people))
	}
}

// 3. DELETE test — state change test
func TestMemoryStore_DeletePerson(t *testing.T) {
	store := &MemoryStore{
		Data: []Person{
			{Name: "Bob", Age: 30, Phone: "111"},
			{Name: "Alice", Age: 25, Phone: "222"},
		},
	}

	store.Data = append(store.Data[:0], store.Data[1:]...)

	if len(store.Data) != 1 {
		t.Fatalf("Expected 1 person after delete, got %d", len(store.Data))
	}

	if store.Data[0].Name != "Ayşe" {
		t.Errorf("Expected Ayşe after delete, got %s", store.Data[0].Name)
	}
}

// 4. UPDATE test — state change test
func TestMemoryStore_UpdatePerson(t *testing.T) {
	store := &MemoryStore{
		Data: []Person{
			{Name: "Ali", Age: 30, Phone: "111"},
		},
	}

	updated := Person{Name: "Mehmet", Age: 40, Phone: "999"}
	store.Data[0] = updated

	if store.Data[0] != updated {
		t.Errorf("Update failed: got %+v, want %+v", store.Data[0], updated)
	}
}
