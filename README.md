# Spatial domain steganography

## Description

**Pseudo-random intervals method** consists in randomly distributing bits of the 
secret message across the container-image.

## Implementation

This program is a Go module that provides message embedding and extracting procedure
to/from image-container. 

To build the source, execute

`go build -o bin/psrand-stego cmd/main.go`

Or just run it w/o building the binary:

`go run cmd/main.go`