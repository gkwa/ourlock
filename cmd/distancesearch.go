/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/gkwa/ourlock/common"
	"github.com/gkwa/ourlock/distancesearch"
	"github.com/spf13/cobra"
)

var chars int

var distancesearchCmd = &cobra.Command{
	Use:     "distancesearch",
	Short:   "Search with all permutations and .{0,chars}+? pattern",
	Aliases: []string{"ds"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Usage: ourlock distancesearch --chars=<chars> <string1> <string2>")
			os.Exit(1)
		}

		distancesearch.Run(args, Directories, chars)
	},
}

func init() {
	rootCmd.AddCommand(distancesearchCmd)
	distancesearchCmd.Flags().IntVar(&chars, "distance", 1000, "maximum number of characters between search terms")
	distancesearchCmd.Flags().BoolVar(&common.DryRun, "dry-run", false, "show the rg command without running it")
}
