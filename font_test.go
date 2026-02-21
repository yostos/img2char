package main

import "testing"

func TestPackFont_AllZeroReturnsZero(t *testing.T) {
	pattern := [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	got := packFont(pattern)
	if got != 0 {
		t.Errorf("packFont(all zero) = 0x%X, want 0x0", got)
	}
}

func TestPackFont_AllFFReturnsAllOnes(t *testing.T) {
	pattern := [8]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	got := packFont(pattern)
	if got != 0xFFFFFFFFFFFFFFFF {
		t.Errorf("packFont(all FF) = 0x%X, want 0xFFFFFFFFFFFFFFFF", got)
	}
}

func TestPackFont_FirstByteOnly(t *testing.T) {
	pattern := [8]byte{0xAB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	got := packFont(pattern)
	if got != 0xAB {
		t.Errorf("packFont(first byte 0xAB) = 0x%X, want 0xAB", got)
	}
}

func TestPackFont_LastByteOnly(t *testing.T) {
	pattern := [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCD}
	got := packFont(pattern)
	want := uint64(0xCD) << 56
	if got != want {
		t.Errorf("packFont(last byte 0xCD) = 0x%X, want 0x%X", got, want)
	}
}

func TestPackFont_ExclamationMarkPattern(t *testing.T) {
	// '!' = font8x8Basic[1]
	got := packFont(font8x8Basic[1])
	// {0x18, 0x3C, 0x3C, 0x18, 0x18, 0x00, 0x18, 0x00}
	want := uint64(0x18) |
		uint64(0x3C)<<8 |
		uint64(0x3C)<<16 |
		uint64(0x18)<<24 |
		uint64(0x18)<<32 |
		uint64(0x00)<<40 |
		uint64(0x18)<<48 |
		uint64(0x00)<<56
	if got != want {
		t.Errorf("packFont('!') = 0x%X, want 0x%X", got, want)
	}
}
