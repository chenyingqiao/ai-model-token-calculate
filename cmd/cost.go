package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"ai-model-token-calculate/tokenizer"
	"ai-model-token-calculate/util"

	"github.com/spf13/cobra"
)

var (
	costFile   string
	costModel  string
	costDir    string
	costJSON   bool
)

func init() {
	cmd := &cobra.Command{
		Use:   "cost",
		Short: "Estimate cost of tokens from file, stdin or directory",
		Run:   runCost,
	}
	cmd.Flags().StringVarP(&costFile, "file", "f", "", "Path to input file")
	cmd.Flags().StringVarP(&costModel, "model", "m", "gpt-4o", "Model name (e.g., gpt-4o, gpt-3.5)")
	cmd.Flags().BoolVar(&costJSON, "json", false, "Output as JSON")
	cmd.Flags().StringVarP(&costDir, "dir", "d", "", "Analyze all .txt files in directory")
	rootCmd.AddCommand(cmd)
}

func runCost(cmd *cobra.Command, args []string) {
	model := strings.ToLower(costModel)
	info, ok := tokenizer.ModelMap[model]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unsupported model: %s", model)
		os.Exit(1)
	}

	type costResult struct {
		File       string  `json:"file,omitempty"`
		TokenCount int     `json:"token_count"`
		CostUSD    float64 `json:"cost_usd"`
	}

	printResult := func(res costResult) {
		if costJSON {
			b, _ := json.Marshal(res)
			fmt.Println(string(b))
		} else {
			if res.File != "" {
				fmt.Printf("[File: %s] ", res.File)
			}
			fmt.Printf("Tokens: %d, Estimated cost: $%.6f", res.TokenCount, res.CostUSD)
		}
	}

	if costDir != "" {
		files, err := util.ReadDirectory(costDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read directory: %v", err)
			os.Exit(1)
		}
		for _, f := range files {
			text, err := util.ReadInput(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read file %s: %v", f, err)
				continue
			}
			count, _, _ := tokenizer.CountTokens(text, info.Tokenizer)
			cost := tokenizer.EstimateCost(count, info.TokenPrice)
			printResult(costResult{File: f, TokenCount: count, CostUSD: cost})
		}
		return
	}

	input, err := util.ReadInput(costFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read error: %v", err)
		os.Exit(1)
	}
	count, _, err := tokenizer.CountTokens(input, info.Tokenizer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Tokenizer error: %v", err)
		os.Exit(1)
	}
	cost := tokenizer.EstimateCost(count, info.TokenPrice)
	printResult(costResult{TokenCount: count, CostUSD: cost})
}
