# ASCII Art Converter — Concept

## Vision

A tool that converts monochrome images to ASCII art based on character shape similarity, rather than pixel density.

## Background

Around 1986, on a Sharp X1 computer, I wrote a BASIC program that read 8x8 dot blocks from the graphic VRAM and matched them against character patterns in the character ROM to automatically convert images into ASCII art. This was motivated by posting to BBS (Bulletin Board System) forums on Japanese personal computer networks — an early attempt at automatic image-to-ASCII-art conversion.

The X1 is long gone, but this project recreates that idea on a modern platform.

## Core Idea

Conventional ASCII art converters represent pixel brightness using character density. This tool takes a different approach: it treats each character as an 8x8 dot shape pattern and selects the most similar character by pattern similarity (Hamming distance) against image blocks.

This produces ASCII art that reflects not only brightness but also line direction and contour shapes.

## Goals

- Faithfully reproduce the original algorithm from that era
- A simple CLI tool with minimal external dependencies
- Prioritize the fun of the idea and reproducibility over practical utility
