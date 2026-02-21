package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestIsWhite_PureWhiteReturnsTrue(t *testing.T) {
	c := color.White
	if !isWhite(c) {
		t.Error("isWhite(White) = false, want true")
	}
}

func TestIsWhite_PureBlackReturnsFalse(t *testing.T) {
	c := color.Black
	if isWhite(c) {
		t.Error("isWhite(Black) = true, want false")
	}
}

func TestIsWhite_MidGrayBoundaryReturnsFalse(t *testing.T) {
	c := color.RGBA{R: 0x7F, G: 0x7F, B: 0x7F, A: 0xFF}
	if isWhite(c) {
		t.Error("isWhite(mid gray) = true, want false")
	}
}

func TestIsWhite_AboveThresholdReturnsTrue(t *testing.T) {
	c := color.RGBA{R: 0x90, G: 0x90, B: 0x90, A: 0xFF}
	if !isWhite(c) {
		t.Error("isWhite(light gray) = false, want true")
	}
}

func TestIsWhite_HighRedOnlyReturnsFalse(t *testing.T) {
	// Only R is high but average luminance stays below threshold
	c := color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
	if isWhite(c) {
		t.Error("isWhite(pure red) = true, want false")
	}
}

func TestLoadAndBlock_640x200(t *testing.T) {
	blocks, cols, rows, err := loadAndBlock("testimages/sample640x200.png")
	if err != nil {
		t.Fatalf("loadAndBlock(640x200): %v", err)
	}
	if cols != 80 {
		t.Errorf("cols = %d, want 80", cols)
	}
	if rows != 25 {
		t.Errorf("rows = %d, want 25", rows)
	}
	if len(blocks) != 25 {
		t.Errorf("len(blocks) = %d, want 25", len(blocks))
	}
	if len(blocks[0]) != 80 {
		t.Errorf("len(blocks[0]) = %d, want 80", len(blocks[0]))
	}
}

func TestLoadAndBlock_320x200(t *testing.T) {
	blocks, cols, rows, err := loadAndBlock("testimages/sample320x200.png")
	if err != nil {
		t.Fatalf("loadAndBlock(320x200): %v", err)
	}
	if cols != 40 {
		t.Errorf("cols = %d, want 40", cols)
	}
	if rows != 25 {
		t.Errorf("rows = %d, want 25", rows)
	}
	if len(blocks) != 25 {
		t.Errorf("len(blocks) = %d, want 25", len(blocks))
	}
	if len(blocks[0]) != 40 {
		t.Errorf("len(blocks[0]) = %d, want 40", len(blocks[0]))
	}
}

func TestLoadAndBlock_UnsupportedResolutionReturnsError(t *testing.T) {
	// Dynamically generate a 100x100 PNG
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.png")
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		t.Fatalf("encode png: %v", err)
	}
	f.Close()

	_, _, _, err = loadAndBlock(path)
	if err == nil {
		t.Fatal("loadAndBlock(100x100) should return error")
	}
	if got := err.Error(); !contains(got, "unsupported resolution") {
		t.Errorf("error = %q, want to contain 'unsupported resolution'", got)
	}
}

func TestLoadAndBlock_NonexistentFileReturnsError(t *testing.T) {
	_, _, _, err := loadAndBlock("nonexistent.png")
	if err == nil {
		t.Fatal("loadAndBlock(nonexistent) should return error")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchSubstring(s, substr)
}

func searchSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
