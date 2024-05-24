package cmd

import (
	"fmt"
	"os"

	"github.com/gkwa/ourlock/common"
	"github.com/gkwa/ourlock/existancesearch"
	"github.com/spf13/cobra"
)

var existancesearchCmd = &cobra.Command{
	Use:     "existancesearch",
	Short:   "Search with all permutations and .+? pattern",
	Aliases: []string{"es"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: ourlock existancesearch <string1> [<string2> ...]")
			os.Exit(1)
		}

		existancesearch.Run(args, common.Directories)
	},
}

func init() {
	rootCmd.AddCommand(existancesearchCmd)
	existancesearchCmd.Flags().BoolVar(&common.DryRun, "dry-run", false, "show the rg command without running it")
}
