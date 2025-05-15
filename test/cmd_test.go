package test

import (
	"testing"

	"ai-model-token-calculate/cmd"
)

func TestRootCommand(t *testing.T) {
	root := cmd.RootCmd()
	if root == nil {
		t.Fatal("RootCmd is nil")
	}
	// 测试子命令是否注册
	foundTokenize := false
	foundCost := false
	foundDecode := false
	for _, c := range root.Commands() {
		switch c.Name() {
		case "tokenize":
			foundTokenize = true
		case "cost":
			foundCost = true
		case "decode":
			foundDecode = true
		}
	}
	if !foundTokenize || !foundCost || !foundDecode {
		t.Fatal("Expected subcommands tokenize, cost, decode registered")
	}
}

