package pkg

import (
	"encoding/binary"
	"image"
	"image/color"
)

// makeEqualToMessageBit makes the least significant bit of specified pixel component equal to current message
// bit. This way the message bit is saved in the image pixel.
func makeEqualToMessageBit(component uint32, messageBit byte) uint32 {
	if component%2 == uint32(messageBit) {
		return component
	} else {
		if messageBit == 1 {
			component++
			return component
		} else {
			component--
			return component
		}
	}
}

// EmbedMessage embeds given message into image-container so this message becomes hidden inside the image. It also returns
// the number of bits which fit inside the container.
func EmbedMessage(img image.Image, message []byte, key []byte) (image.Image, int) {

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	imgLen := width * height

	newImg := ConvertToRGBA(img)

	var messageIdx, keyIdx int

	for i := int(key[keyIdx]); i < imgLen; i += int(key[keyIdx]) {
		if messageIdx < len(message) {

			x := i % width
			y := i / width

			pix := newImg.At(x, y)
			r, g, b, a := pix.RGBA()

			b = makeEqualToMessageBit(b, message[messageIdx])

			newPix := color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a),
			}

			newImg.Set(x, y, newPix)
		} else {
			break
		}

		messageIdx++
		keyIdx = messageIdx % len(key)
	}

	return newImg, messageIdx
}

// ExtractMessage extracts hidden message from image-container and returns it.
func ExtractMessage(img image.Image, messageLen int, key []byte) []byte {

	bounds := img.Bounds()
	width := bounds.Max.X

	var message []byte
	keyIdx := int(key[0])
	imgIdx := keyIdx

	for messageIdx := 0; messageIdx < messageLen; {

		x := imgIdx % width
		y := imgIdx / width

		pix := img.At(x, y)
		_, _, b, _ := pix.RGBA()

		if b%2 == 1 {
			message = append(message, 1)
		} else {
			message = append(message, 0)
		}

		messageIdx++
		keyIdx = messageIdx % len(key)
		imgIdx += int(key[keyIdx])
	}

	return message
}

// EmbedMetadata embeds random intervals' key and message length to the end of the file.
func EmbedMetadata(file []byte, messageLen int, key []byte) []byte {
	ml := make([]byte, 4)
	binary.LittleEndian.PutUint32(ml, uint32(messageLen))

	file = append(file, 0x55)
	file = append(file, ml...)
	file = append(file, key...)
	file = append(file, 0x55)

	return file
}

// DetectEmbedding checks the end of given file on presence of random intervals' key and message length embedded before
// between special separators. Returns the detection flag.
func DetectEmbedding(file []byte, keyLen int) bool {
	if file[len(file)-1] == 0x55 && file[len(file)-6-keyLen] == 0x55 {
		return true
	}
	return false
}

// ExtractMetadata extracts random intervals' key and hidden message length from the end of the file.
func ExtractMetadata(file []byte, keyLen int) (int, []byte) {
	key := file[len(file)-keyLen-1 : len(file)-1]

	ml := file[len(file)-keyLen-4-1 : len(file)-keyLen-1]
	messageLen := binary.LittleEndian.Uint32(ml)

	return int(messageLen), key
}
