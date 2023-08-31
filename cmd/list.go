package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/spf13/cobra"
)

const notesPerPage = 5

type Note struct {
	Title        string
	Directory    string // Added Directory for greater granularity
	Path         string
	DateCreated  time.Time
	DateModified time.Time
}

func fetchNotes(category string) ([]Note, error) { // Added error to return type
	basePath, err := getBasePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get base path: %w", err)
	}

	searchDir := basePath
	if category != "" {
		searchDir = filepath.Join(basePath, category)
	}

	var notes []Note
	err = filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".md" {
			creationTime := info.ModTime()
			notes = append(notes, Note{
				Title:        filepath.Base(path),
				Directory:    filepath.Dir(path), // Added Directory
				Path:         path,
				DateCreated:  creationTime,
				DateModified: creationTime,
			})
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %v: %w", searchDir, err)
	}
	return notes, nil
}

func listNotes(category string) {
	notes, err := fetchNotes(category)
	if err != nil {
		fmt.Printf("Error fetching notes: %s\n", err)
		return
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[i].DateModified.After(notes[j].DateModified)
	})

	for i := 0; i < notesPerPage && i < len(notes); i++ {
		note := notes[i]
		fmt.Printf("Title: %s\nDate Created: %s\nDate Modified: %s\n\n",
			note.Title, note.DateCreated.Format("2006-01-02"), note.DateModified.Format("2006-01-02"))
	}
}

func listNotesWithSummary(category string) {
	// TODO: Implement logic to display note summaries
}

var listCmd = &cobra.Command{
	Use:   "list [summary]",
	Short: "List notes with optional summary",
	Long:  "Displays a list of notes, with optional summary of each note",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		if len(args) > 0 && args[0] == "summary" {
			listNotesWithSummary(category)
		} else {
			listNotes(category)
		}
	},
}

func init() {
	listCmd.Flags().String("category", "", "Specify category to list notes from")
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(listCategoriesCmd)
}
