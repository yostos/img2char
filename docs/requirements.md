# ASCII Art Converter — Requirements

## Command Line

```
img2char [-v] <image-file>
```

### Arguments

| Argument | Description |
|----------|-------------|
| `image-file` | Path to the input image file (required) |

### Options

| Option | Description | Default |
|--------|-------------|---------|
| `-v` | Print processing info to stderr | off |

## Input

- Format: PNG only
- Content: Pre-binarized monochrome image (binarization is outside the scope of this tool)
- Resolution: Only 640x200 or 320x200 are accepted; any other resolution results in an error

## Output

- Print ASCII text to stdout
- One block (8x8 pixels) = one character
- Character set: Printable ASCII (0x20–0x7E, 95 characters)
- 640x200 → 80x25 characters, 320x200 → 40x25 characters
- A newline is appended at the end of each row
- File saving is handled via shell redirection
