package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "dev"

// Main Routine
func main() {
	verbose := flag.Bool("v", false, "print processing info to stderr")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-v] [-version] <image-file>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	imagePath := flag.Arg(0)

	// Precompute font table as uint64 values.
	var fontTable [95]uint64
	for i := 0; i < 95; i++ {
		fontTable[i] = packFont(font8x8Basic[i])
	}

	// Load image and split into 8x8 blocks.
	blocks, cols, rows, err := loadAndBlock(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Fprintf(os.Stderr, "image: %s\n", imagePath)
		fmt.Fprintf(os.Stderr, "blocks: %dx%d (%d total)\n", cols, rows, cols*rows)
	}

	// Match each block and output characters.
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			ch := matchChar(blocks[r][c], fontTable)
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}
