package cmd

import (
	"os"
	"testing"
	"time"
)

func TestGetBasePath(t *testing.T) {
	os.Setenv("GOTES_PATH", "/path/from/env")
	path, err := getBasePath()
	if err != nil || path != "/path/from/env" {
		t.Fatalf("Expected /path/from/env but got %s", path)
	}

	os.Unsetenv("GOTES_PATH")
	path, err = getBasePath()
	home, _ := os.UserHomeDir()
	if err != nil || path != home {
		t.Fatalf("Expected home directory but got %s", path)
	}
}

func TestEnsureDirExists(t *testing.T) {
	tmpDir := os.TempDir()

	testDir := tmpDir + "/testDir"
	os.RemoveAll(testDir) // Make sure it doesn't exist

	if err := createDirectory(testDir); err != nil {
		t.Fatalf("Error creating directory: %s", err)
	}

	// Check if directory was actually created
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Fatalf("Directory was not created")
	}

	// Try creating the directory again; it should not error since it exists
	if err := createDirectory(testDir); err != nil {
		t.Fatalf("Error when directory exists: %s", err)
	}
}

func TestGetCurrentDate(t *testing.T) {
	expected := time.Now().Format("2006-01-02")
	got := getCurrentDate()
	if got != expected {
		t.Errorf("Expected %s but got %s", expected, got)
	}
}
