package main

import (
	"os"
	"path/filepath"
	"testing"
)

// example/excel/sample.xlsx を変換した結果が
// example/build/sample/sheet1.md (ゴールデン) と一致することを確認する
func TestReadWriteSheet_Golden(t *testing.T) {
	outputDir := t.TempDir()

	if err := readWriteSheet(filepath.Join("example", "excel", "sample.xlsx"), outputDir); err != nil {
		t.Fatalf("readWriteSheet: %v", err)
	}

	got, err := os.ReadFile(filepath.Join(outputDir, "sheet1.md"))
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	want, err := os.ReadFile(filepath.Join("example", "build", "sample", "sheet1.md"))
	if err != nil {
		t.Fatalf("read golden: %v", err)
	}

	if string(got) != string(want) {
		t.Errorf("output mismatch with golden file\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}
