package cmd

import (
	"ai-model-token-calculate/tokenizer"
	"ai-model-token-calculate/util"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	tokenizeFile    string
	tokenizeModel   string
	tokenizeVerbose bool
	tokenizeJSON    bool
	tokenizeDir     string
)

func init() {
	cmd := &cobra.Command{
		Use:   "tokenize",
		Short: "Count tokens from file, stdin or directory",
		Run:   runTokenize,
	}
	cmd.Flags().StringVarP(&tokenizeFile, "file", "f", "", "Path to input file")
	cmd.Flags().StringVarP(&tokenizeModel, "model", "m", "gpt-4o", "Model name (e.g., gpt-4o, gpt-3.5)")
	cmd.Flags().BoolVarP(&tokenizeVerbose, "verbose", "v", false, "Print token IDs")
	cmd.Flags().BoolVar(&tokenizeJSON, "json", false, "Output as JSON")
	cmd.Flags().StringVarP(&tokenizeDir, "dir", "d", "", "Analyze all .txt files in directory")
	rootCmd.AddCommand(cmd)
}

func runTokenize(cmd *cobra.Command, args []string) {
	model := strings.ToLower(tokenizeModel)
	info, ok := tokenizer.ModelMap[model]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unsupported model: %s", model)
		os.Exit(1)
	}

	// Batch mode: Directory
	if tokenizeDir != "" {
		files, err := util.ReadDirectory(tokenizeDir)
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
			count, tokens, _ := tokenizer.CountTokens(text, info.Tokenizer)
			fmt.Printf("[File: %s] Tokens: %d", f, count)
			if tokenizeVerbose {
				fmt.Println(tokens)
			}
		}
		return
	}

	// Single file or stdin
	input, err := util.ReadInput(tokenizeFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read error: %v", err)
		os.Exit(1)
	}
	count, tokens, err := tokenizer.CountTokens(input, info.Tokenizer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Tokenizer error: %v", err)
		os.Exit(1)
	}

	out := tokenizer.Output{
		Model:         model,
		Tokenizer:     info.Tokenizer,
		ContextLength: info.ContextLength,
		TokenCount:    count,
	}
	if tokenizeVerbose {
		out.Tokens = tokens
	}

	if tokenizeJSON {
		json.NewEncoder(os.Stdout).Encode(out)
	} else {
		fmt.Printf("Model: %s (Tokenizer: %s)", model, info.Tokenizer)
		fmt.Printf("Tokens: %d", count)
		if tokenizeVerbose {
			fmt.Println("Token IDs:", tokens)
		}
	}
}
