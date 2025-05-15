package tokenizer

type ModelInfo struct {
	Tokenizer     string
	ContextLength int
	TokenPrice    float64 // USD per token
}

var ModelMap = map[string]ModelInfo{
	"gpt-3.5":     {"cl100k_base", 16384, 0.000002},
	"gpt-4":       {"cl100k_base", 128000, 0.000010},
	"gpt-4o":      {"cl100k_base", 128000, 0.000005},
	"gpt-4o-mini": {"cl100k_base", 128000, 0.000004},
	"davinci":     {"p50k_base", 4097, 0.000020},
}
