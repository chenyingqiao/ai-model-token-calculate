package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"strconv"
	"ai-model-token-calculate/tokenizer"

	"github.com/spf13/cobra"
)

var (
	decodeTokens  string
	decodeModel   string
	decodeJSON    bool
)

func init() {
	cmd := &cobra.Command{
		Use:   "decode",
		Short: "Decode token IDs to string",
		Run:   runDecode,
	}
	cmd.Flags().StringVarP(&decodeTokens, "tokens", "t", "", "Comma separated token IDs")
	cmd.Flags().StringVarP(&decodeModel, "model", "m", "gpt-4o", "Model name (e.g., gpt-4o, gpt-3.5)")
	cmd.Flags().BoolVar(&decodeJSON, "json", false, "Output as JSON")
	rootCmd.AddCommand(cmd)
}

func runDecode(cmd *cobra.Command, args []string) {
	if decodeTokens == "" {
		fmt.Fprintln(os.Stderr, "Please provide token IDs with -t flag")
		os.Exit(1)
	}
	model := strings.ToLower(decodeModel)
	info, ok := tokenizer.ModelMap[model]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unsupported model: %s", model)
		os.Exit(1)
	}

	tokenStrs := strings.Split(decodeTokens, ",")
	var tokens []int
	for _, ts := range tokenStrs {
		t, err := strconv.Atoi(strings.TrimSpace(ts))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid token ID: %s", ts)
			os.Exit(1)
		}
		tokens = append(tokens, t)
	}

	text := tokenizer.DecodeTokens(tokens, info.Tokenizer)
	if decodeJSON {
		type decodeOutput struct {
			Text string `json:"text"`
		}
		out := decodeOutput{Text: text}
		json.NewEncoder(os.Stdout).Encode(out)
	} else {
		fmt.Println(text)
	}
}
