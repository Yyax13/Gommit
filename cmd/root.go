package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "gommit",
    Short: "Gommit - AI-assisted git commit generator",
    Long:  "Gommit helps you generate git commit messages using AI prompts and your git diff history.",
}

// Execute runs the root command
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        panic(err)
    }
}