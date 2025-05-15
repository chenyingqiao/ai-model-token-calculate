package tokenizer

import "github.com/pkoukk/tiktoken-go"

type Output struct {
	Model         string `json:"model"`
	Tokenizer     string `json:"tokenizer"`
	ContextLength int    `json:"context_length"`
	TokenCount    int    `json:"token_count"`
	Tokens        []int  `json:"tokens,omitempty"`
}

func CountTokens(text string, encodingName string) (int, []int, error) {
	encoder, err := tiktoken.GetEncoding(encodingName)
	if err != nil {
		return 0, nil, err
	}
	tokens := encoder.Encode(text, nil, nil)
	return len(tokens), tokens, nil
}

func DecodeTokens(tokens []int, encodingName string) string {
	encoder, err := tiktoken.GetEncoding(encodingName)
	if err != nil {
		return ""
	}
	return encoder.Decode(tokens)
}

func EstimateCost(tokenCount int, price float64) float64 {
	return float64(tokenCount) * price
}
