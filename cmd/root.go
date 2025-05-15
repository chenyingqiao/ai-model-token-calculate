package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ai-model-token-calculate",
	Short: "Token counting and cost estimation tool for AI models",
	Long:  "A CLI tool to count, decode and estimate cost of tokens for various AI models using tiktoken.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
