package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the base command for glyphra
var RootCmd = &cobra.Command{
	Use:   "glyphra",
	Short: "Glyphra is an AI-powered tool that explains infrastructure changes",
	Long: `Glyphra helps engineers understand infrastructure artifacts
like Terraform plans and GitHub PRs using natural language summaries.`,
}

// Execute is the entry point called from main.go
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
