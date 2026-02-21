package main

import "testing"

func TestPackBlock_AllZero(t *testing.T) {
	block := [8]byte{}
	got := packBlock(block)
	if got != 0 {
		t.Errorf("packBlock(all zero) = 0x%X, want 0x0", got)
	}
}

func TestPackBlock_AllFF(t *testing.T) {
	block := [8]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	got := packBlock(block)
	if got != 0xFFFFFFFFFFFFFFFF {
		t.Errorf("packBlock(all FF) = 0x%X, want 0xFFFFFFFFFFFFFFFF", got)
	}
}

func TestPackBlock_SingleByte(t *testing.T) {
	block := [8]byte{0x00, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00}
	got := packBlock(block)
	want := uint64(0x42) << 16
	if got != want {
		t.Errorf("packBlock(byte[2]=0x42) = 0x%X, want 0x%X", got, want)
	}
}

func TestPackBlockAndPackFontConsistency(t *testing.T) {
	// packBlock and packFont share the same logic, so identical input must produce identical output
	input := [8]byte{0x18, 0x3C, 0x3C, 0x18, 0x18, 0x00, 0x18, 0x00}
	fromBlock := packBlock(input)
	fromFont := packFont(input)
	if fromBlock != fromFont {
		t.Errorf("packBlock and packFont disagree: block=0x%X, font=0x%X", fromBlock, fromFont)
	}
}

func TestMatchChar_AllZeroBlockReturnsSpace(t *testing.T) {
	var fontTable [95]uint64
	for i := 0; i < 95; i++ {
		fontTable[i] = packFont(font8x8Basic[i])
	}

	block := [8]byte{}
	got := matchChar(block, fontTable)
	if got != ' ' {
		t.Errorf("matchChar(all zero) = %q, want ' '", got)
	}
}

func TestMatchChar_ExactFontPatternMatch(t *testing.T) {
	var fontTable [95]uint64
	for i := 0; i < 95; i++ {
		fontTable[i] = packFont(font8x8Basic[i])
	}

	// Feed the '!' font pattern directly as a block
	block := font8x8Basic[1] // '!'
	got := matchChar(block, fontTable)
	if got != '!' {
		t.Errorf("matchChar('!' pattern) = %q, want '!'", got)
	}
}

func TestMatchChar_All95CharsRoundTrip(t *testing.T) {
	var fontTable [95]uint64
	for i := 0; i < 95; i++ {
		fontTable[i] = packFont(font8x8Basic[i])
	}

	for i := 0; i < 95; i++ {
		expected := byte(0x20 + i)
		got := matchChar(font8x8Basic[i], fontTable)
		if got != expected {
			t.Errorf("round-trip failed: index=%d, expected=%q (0x%02X), got=%q (0x%02X)",
				i, expected, expected, got, got)
		}
	}
}
