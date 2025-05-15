package test

import (
	"os"
	"path/filepath"
	"testing"

	"ai-model-token-calculate/util"
)

func TestReadInputAndReadDirectory(t *testing.T) {
	// 创建临时文件夹和文件
	dir := t.TempDir()
	file := filepath.Join(dir, "test.txt")
	content := "hello test\n"
	err := os.WriteFile(file, []byte(content), 0644)
	if err != nil {
		t.Fatalf("WriteFile error: %v", err)
	}

	// 测试 ReadInput 读取文件
	got, err := util.ReadInput(file)
	if err != nil {
		t.Fatalf("ReadInput error: %v", err)
	}
	if got != content {
		t.Fatalf("ReadInput got %q, want %q", got, content)
	}

	// 测试 ReadDirectory
	files, err := util.ReadDirectory(dir)
	if err != nil {
		t.Fatalf("ReadDirectory error: %v", err)
	}
	found := false
	for _, f := range files {
		if f == file {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("ReadDirectory did not find test.txt")
	}
}

