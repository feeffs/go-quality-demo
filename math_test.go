package main

import (
	"os"
	"testing"
)

const filePermissions = 0644

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Error("expected error for division by zero")
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := divide(100, 3); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func TestReadFileNotFound(t *testing.T) {
	_, err := readTextFile("nonexistent.txt")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestDivideSuccess(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestReadFileSuccess(t *testing.T) {
	// создаём временный файл
	err := os.WriteFile("test_temp.txt", []byte("hello"), filePermissions)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer func() {
		if err := os.Remove("test_temp.txt"); err != nil {
			t.Errorf("failed to remove temp file: %v", err)
		}
	}() // удаляем после теста

	content, err := readTextFile("test_temp.txt")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if content != "hello" {
		t.Errorf("expected 'hello', got '%s'", content)
	}
}
