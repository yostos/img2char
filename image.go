package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
)

// loadAndBlock reads a PNG file, validates its resolution, and splits it into 8x8 blocks.
// Returns a 2D slice of blocks [row][col], where each block is [8]byte.
func loadAndBlock(path string) ([][][8]byte, int, int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("open image: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("decode image: %w", err)
	}

	bounds := img.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y

	if !((w == 640 && h == 200) || (w == 320 && h == 200)) {
		return nil, 0, 0, fmt.Errorf("unsupported resolution %dx%d (must be 640x200 or 320x200)", w, h)
	}

	cols := w / 8
	rows := h / 8

	blocks := make([][][8]byte, rows)
	for r := 0; r < rows; r++ {
		blocks[r] = make([][8]byte, cols)
		for c := 0; c < cols; c++ {
			var block [8]byte
			for y := 0; y < 8; y++ {
				var row byte
				for x := 0; x < 8; x++ {
					px := img.At(bounds.Min.X+c*8+x, bounds.Min.Y+r*8+y)
					if isWhite(px) {
						row |= 1 << uint(x)
					}
				}
				block[y] = row
			}
			blocks[r][c] = block
		}
	}

	return blocks, cols, rows, nil
}

// isWhite returns true if the pixel is considered white (foreground).
// For a binarized image, we check if the luminance exceeds the midpoint.
func isWhite(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	// RGBA returns values in [0, 0xFFFF]. Use simple average for luminance.
	lum := (r + g + b) / 3
	return lum > 0x7FFF
}
