/*
Copyright © 2024 Eduardo Lopes <eduardotorreslopes@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	files []string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var style = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000")).
			Background(lipgloss.Color("#FFFFFF")).
			Padding(1, 2)

		fmt.Println(style.Render("Git Commiter"))

		unstagedFiles, err := exec.Command("git", "ls-files", "--others", "--exclude-standard").Output()
		if err != nil {
			log.Fatalf("Failed to list unstaged files: %v", err)
		}
		files = strings.Split(strings.TrimSpace(string(unstagedFiles)), "\n")

		newOptions := make([]huh.Option[string], len(files))
		for i, file := range files {
			newOptions[i] = huh.NewOption(file, file).Selected(false)
		}

		stagedFiles, err := exec.Command("git", "ls-files", "--stage").Output()
		if err != nil {
			log.Fatalf("Failed to list staged files: %v", err)
		}

		fmt.Println("Staged files: ", string(stagedFiles))

		// form := huh.NewForm(
		// 	huh.NewGroup(
		// 		huh.NewMultiSelect[string]().
		// 			Title("Select unstaged files").
		// 			Options(newOptions...).
		// 			Value(&files),
		// 	),
		// )

		// err = form.Run()
		// if err != nil {
		// 	log.Fatal(err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
