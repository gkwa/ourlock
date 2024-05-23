package cmd

import (
	"fmt"
	"os"

	"github.com/gkwa/ourlock/core"
	"github.com/spf13/cobra"
)

var (
	dryRun      bool
	directories []string
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Search for strings in specified directories",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: ourlock search <string1> [<string2> ...]")
			os.Exit(1)
		}
		if len(directories) == 0 {
			directories = []string{
				"/Users/mtm/Documents/Obsidian Vault/",
				"~/pdev/taylormonacelli/notes/",
			}
		}
		core.Run(args, directories, dryRun)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVar(&dryRun, "dry-run", false, "show the rg command without running it")
	searchCmd.Flags().StringArrayVar(&directories, "dir", []string{}, "directories to search in (default: predefined directories)")
}
