package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const NotesPerPage = 5 // Capitalized to indicate package-wide constant

type Note struct {
	Title        string
	Directory    string // Make sure this is used in the future.
	Path         string
	DateCreated  time.Time
	DateModified time.Time
}

func fetchNotes(category string) ([]Note, error) {
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
				Directory:    filepath.Dir(path),
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

func listNotesInteractive(category string) {
	notes, err := fetchNotes(category)
	if err != nil {
		fmt.Printf("Error fetching notes: %s\n", err)
		return
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[i].DateModified.After(notes[j].DateModified)
	})

	// Prepare the notes for the prompt
	noteTitles := []string{}
	notePaths := []string{}
	for _, note := range notes {
		displayText := fmt.Sprintf("[%s] %s - %s", note.Directory, note.Title, note.DateCreated.Format("2006-01-02"))
		noteTitles = append(noteTitles, displayText)
		notePaths = append(notePaths, note.Path)
	}

	prompt := promptui.Select{
		Label: "Select Note",
		Items: noteTitles,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Open the selected note in vim
	cmd := exec.Command("vim", notePaths[index])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to open note in vim: %s\n", err)
		return
	}
}

func listNotesWithSummary(w io.Writer, category string) {
	// TODO: Implement logic to display note summaries
}

var listCmd = &cobra.Command{
	Use:   "list [summary]",
	Short: "List notes with optional summary",
	Long:  "Displays a list of notes, with optional summary of each note",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		category, err := cmd.Flags().GetString("category")
		if err != nil {
			fmt.Printf("Error getting category: %s\n", err)
			return
		}
		if len(args) > 0 && args[0] == "summary" {
			listNotesWithSummary(os.Stdout, category)
		} else {
			listNotesInteractive(category)
		}
	},
}

func init() {
	listCmd.Flags().String("category", "", "Specify category to list notes from")
	rootCmd.AddCommand(listCmd)
}
