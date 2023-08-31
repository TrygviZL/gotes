package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func getBasePath() (string, error) {
	basePath := os.Getenv("GOTES_PATH")
	if basePath == "" {
		return os.UserHomeDir()
	}
	return basePath, nil
}

func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755) // 0755 is for permissions
	}
	return nil
}

var createCmd = &cobra.Command{
	Use:   "new <category> <title>",
	Short: "Create a new markdown note",
	Long:  "Create a new note within the given category and title",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		category := args[0]
		title := args[1]

		basePath, err := getBasePath()
		if err != nil {
			fmt.Println("Error retrieving home directory:", err)
			return
		}

		// Create the directory if it does not exist
		fullPath := filepath.Join(basePath, category, title+".md")
		if err := createDirectory(filepath.Dir(fullPath)); err != nil {
			fmt.Println("Error creating directories:", err)
			return
		}

		// create the markdown file
		file, err := os.Create(fullPath)
		if err != nil {
			fmt.Println("error creating markdown file:", err)
			return
		}
		defer file.Close()

		fmt.Printf("Note created at %s\n", fullPath)

		// Add metadata to file
		currentDate := getCurrentDate()
		frontMatter := fmt.Sprintf(`---
		title: "%s"
		category: "%s"
		date_created: %s
		date_modified: %s
		---

		`, title, category, currentDate, currentDate)

		// write frontMatter to file before opening in vim
		_, err = file.WriteString(frontMatter)
		if err != nil {
			fmt.Println("error writing metadata to markdown file:", err)
			os.Exit(1)
		}

		// open file
		vimCmd := exec.Command("vim", fullPath)
		vimCmd.Stdin = os.Stdin
		vimCmd.Stdout = os.Stdout
		vimCmd.Stderr = os.Stderr
		err = vimCmd.Run()
		if err != nil {
			fmt.Println("Error opening file in vim:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
