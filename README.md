ai-model-token-calculate
=========================

A command-line tool written in Go for analyzing OpenAI token usage, including:

📘 中文版本说明请点击：[简体中文 README](README.zh.md)


- Token counting for input text or files
- Decoding token sequences back into strings
- Estimating token cost based on selected models
- Batch analysis for directories containing text files
- Extensible subcommand structure using Cobra

----------------------------------------
Features
----------------------------------------

Commands:

  tokenize   - Count tokens from text or file input
  decode     - Decode a sequence of tokens back to string
  cost       - Estimate cost based on token count and model price
  analyze    - Analyze all `.txt` files in a directory (batch tokenize + cost)

----------------------------------------
Installation & Usage
----------------------------------------

1. Clone the project:

   git clone https://github.com/yourname/ai-model-token-calculate.git
   cd ai-model-token-calculate

2. Install dependencies:

   go mod tidy

3. Build the CLI tool:

   go build -o tokencli

----------------------------------------
Examples
----------------------------------------

1. Tokenize a file:

   ./tokencli tokenize -f input.txt -m gpt-4o -v

2. Tokenize from standard input:

   echo "Hello world!" | ./tokencli tokenize -m gpt-3.5

3. Decode tokens:

   ./tokencli decode -t "1212,402,98" -m gpt-4

4. Estimate cost:

   ./tokencli cost -f input.txt -m gpt-4o

5. Analyze all `.txt` files in a directory:

   ./tokencli analyze -d ./docs -m gpt-3.5

----------------------------------------
Supported Models and Prices (per 1,000 tokens)
----------------------------------------

Model:         gpt-4o
Encoder:       cl100k_base
Input Price:   $0.000005
Output Price:  $0.000015

Model:         gpt-4
Encoder:       cl100k_base
Input Price:   $0.00003
Output Price:  $0.00006

Model:         gpt-3.5-turbo
Encoder:       cl100k_base
Input Price:   $0.0005
Output Price:  $0.0015

----------------------------------------
Testing
----------------------------------------

To run all unit tests:

   go test ./test/...

----------------------------------------
Project Structure
----------------------------------------

.
├── cmd          - CLI command implementations
├── tokenizer    - Token encoding/decoding logic
├── util         - File reading and utility functions
├── test         - Unit tests
├── main.go
└── README.txt

----------------------------------------
License
----------------------------------------

MIT
