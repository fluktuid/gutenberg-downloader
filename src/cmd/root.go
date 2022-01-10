package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gutenber-downloader",
	Short: "gutenber-downloader is a golang-based downloader of gutenberg site content",
	Long: `A Downloader for your free and open Gutenberg books.
	              If you want more free sites create an issue at
								https://github.com/fluktuid/gutenberg-downloader`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
