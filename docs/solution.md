# ASCII Art Converter — Solution Design

## Font Pattern Generation

- Use font8x8 data
  - Source: https://github.com/dhepper/font8x8
  - License: Public Domain
- Embed 8x8 bitmaps for 95 printable ASCII characters (0x20–0x7E) in Go source
- Each character is represented as 8 bytes (one byte per row)

## Image Loading & Blocking

- Load the PNG image
- Validate resolution (error if not 640x200 or 320x200)
- Split into 8x8 pixel blocks from top-left, scanning left to right
- Convert each block to the same format as the font patterns (8 bytes)
  - 640x200 → 80x25 = 2000 blocks
  - 320x200 → 40x25 = 1000 blocks

## Matching (Per-Block Loop)

For each block, repeat the following:

1. XOR the block (8 bytes) with the font pattern (8 bytes)
2. Count the set bits (popcount = Hamming distance)
3. Compare against all 95 characters and select the one with the minimum Hamming distance
4. Write the selected character to stdout
5. Output a newline at the end of each row (after the 80th or 40th character)

## Go Standard Library Usage

- `image/png` — PNG decoding
- `math/bits.OnesCount64` — popcount
- `os`, `fmt` — I/O
