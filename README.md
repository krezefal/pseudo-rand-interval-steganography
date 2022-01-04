# Spatial domain steganography

## Description

The **pseudo-random interval method** involves changing the components of certain 
image pixels through a pseudorandom interval set by a function or a secret key. 
This way, the watermark bits are distributed over the entire array of pixels of 
the image-container.

The last bit of the selected pixel component (blue in current realisation) becomes
equal to the current bit of the message (one bit is subtracted or added from/to
the component, if it needed). The image containing the watermark is visually 
indistinguishable from the original one.

This method is especially effective when the watermark bit length is 
significantly less than the number of pixels in image.

## Implementation

The program is a CLI application that performs watermark embedding and extracting
procedure to/from image-container. This program is a Go module.

To build the binary, execute

`go build -o bin/psrand-stego cmd/main.go`

Or just run it w/o building the binary:

`go run cmd/main.go`

Folder with examples contains a couple of files on which program can be tested 
"out of box".

### Using

To run the program in embedding mode use next example:

`<psrand-stego> -src <path/to/src.bmp> -m <binary message> -tg <path/to/tg.bmp>`

To run the program in extracting mode:

`<psrand-stego> -src <path/to/src.bmp> -ext`

where\
`<psrand-stego>` is the name of the binary for particular platform\
`-src` is the flag followed by the path to source file\
`-ext` means extract. This flag specified only if extraction procedure is 
required\
`-m` is the flag followed by the binary sequence (watermark)\
`-tg` is the flag followed by the path to target file

For example:\
`./psrand-stego -src ../examples/ford.bmp -m 10101010 -tg 
../examples/poison_ford.bmp`\
or\
`./psrand-stego -src ../examples/poison_ford.bmp -ext`

It is also available to embed a randomly generated binary sequence using `-r` 
(info about this and other optional flags is provided with `-h` option).

### Principle of operation

*add principle of operation*

## Theory. Research

*add theory basics and psnr observations*