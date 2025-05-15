package test

import (
	"ai-model-token-calculate/tokenizer"
	"testing"
)

func TestCountAndDecodeTokens(t *testing.T) {
	text := "Hello, world!"
	count, tokens, err := tokenizer.CountTokens(text, "cl100k_base")
	if err != nil {
		t.Fatalf("CountTokens failed: %v", err)
	}
	if count == 0 {
		t.Fatal("Expected token count > 0")
	}

	decoded := tokenizer.DecodeTokens(tokens, "cl100k_base")
	if decoded == "" {
		t.Fatal("Decoded text is empty")
	}
}
