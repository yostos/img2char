# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

img2char is a Go CLI tool that converts monochrome binary PNG images to ASCII art by matching 8x8 pixel blocks against font character patterns using Hamming distance (XOR + popcount). Unlike conventional ASCII art converters that map pixel density, this tool selects characters based on shape similarity.

## Build & Run

```bash
go build -o img2char          # Build binary
./img2char [-v] <image.png>   # Run (outputs ASCII to stdout)
go test -v ./...              # Run tests
go vet ./...                  # Static analysis
```

The `convert.sh` wrapper script preprocesses arbitrary images via ImageMagick (`magick`) before feeding them to img2char:
```bash
./convert.sh [-w 640|320] [-t threshold] [-n] <input-image>
```

## Architecture

Three-stage pipeline, all in package `main` with zero external dependencies:

1. **`font.go`** — `font8x8Basic`: embedded 8x8 bitmap data for 95 printable ASCII characters (0x20–0x7E, from [font8x8](https://github.com/dhepper/font8x8)). `packFont()` converts `[8]byte` to `uint64` for fast comparison.

2. **`image.go`** — `loadAndBlock()`: reads PNG, validates resolution (must be exactly 640x200 or 320x200), splits into 8x8 blocks as `[8]byte`. `isWhite()` binarizes pixels by luminance threshold.

3. **`match.go`** — `matchChar()`: packs block to `uint64`, XORs against all 95 font patterns, selects character with minimum `bits.OnesCount64` (Hamming distance).

`main.go` orchestrates: parse flags → build font table → load/block image → match each block → print character grid (80x25 or 40x25).

## Constraints

- Input must be pre-binarized PNG at exactly 640x200 or 320x200
- Output is plain ASCII text to stdout; file saving via shell redirection
- No go.sum or external modules — standard library only
