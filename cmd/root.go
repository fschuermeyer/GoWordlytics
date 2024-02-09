/*
Copyright © 2024 Felix Schürmeyer
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fschuermeyer/GoWordlytics/internal/analyze"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GoWordlytics",
	Short: "GoWordlytics is a simple CLI tool to analyze a WordPress Website for Plugins, Themes, and other details.",
	Long:  `GoWordlytics is a CLI tool to analyze a WordPress Website for Plugins, Themes, and other details. It is a simple tool to get an overview of a WordPress Website. It is written in Go and uses the Cobra library for the CLI.`,

	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")

		if url == "" {
			fmt.Println("Please provide a URL to analyze")
			os.Exit(1)
		}

		report, err := analyze.NewReport(url)

		if err == analyze.ERR_MALFORMED_URL {
			fmt.Println("The URL is malformed")
			os.Exit(1)
		}

		report.Output()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "The URL of the WordPress Website to analyze.")
}
