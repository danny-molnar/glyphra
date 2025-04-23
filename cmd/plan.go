package cmd

import (
	"fmt"
	"os"

	"github.com/danny-molnar/glyphra/pkg/llm"
	"github.com/spf13/cobra"
)

var inputFile string

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Parse and explain a Terraform plan",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Println("❌ Failed to read input file:", err)
			return
		}

		client := llm.NewHuggingFaceClient("google/flan-t5-large")
		output, err := client.Explain("Summarize the following Terraform plan:\n\n" + string(content))
		if err != nil {
			fmt.Println("💥 AI failed:", err)
			return
		}

		fmt.Println("🧠 Explanation:\n")
		fmt.Println(output)
	},
}

func init() {
	planCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to Terraform plan file")
	planCmd.MarkFlagRequired("input")
	RootCmd.AddCommand(planCmd)
}
