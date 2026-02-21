package main

import "math/bits"

// packBlock packs an 8-byte block into a uint64.
func packBlock(block [8]byte) uint64 {
	var v uint64
	for i := 0; i < 8; i++ {
		v |= uint64(block[i]) << (i * 8)
	}
	return v
}

// matchChar finds the printable ASCII character whose font pattern is closest
// to the given block, measured by Hamming distance.
func matchChar(block [8]byte, fontTable [95]uint64) byte {
	blockVal := packBlock(block)
	bestDist := 65 // max possible is 64
	bestIdx := 0
	for i := 0; i < 95; i++ {
		dist := bits.OnesCount64(blockVal ^ fontTable[i])
		if dist < bestDist {
			bestDist = dist
			bestIdx = i
		}
	}
	return byte(0x20 + bestIdx)
}
