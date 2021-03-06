package pkg

import (
	"golang.org/x/image/bmp"
	"image"
	"math/rand"
	"os"
	"time"
)

// ReadImage reads image at specified path. For bmp only.
func ReadImage(filePath string) (image.Image, error) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return bmp.Decode(imgFile)
}

// WriteImage writes image to specified path. For bmp only.
func WriteImage(filePath string, img image.Image) error {
	if file, err := os.Create(filePath); err != nil {
		return err
	} else {
		if imgErr := bmp.Encode(file, img); imgErr != nil {
			return imgErr
		}
	}

	return nil
}

// ConvertToRGBA translates specified instance of image represented as Image interface into RGBA structure in order to
// be able to modify image array later.
func ConvertToRGBA(img image.Image) *image.RGBA {

	bounds := img.Bounds()
	imgRGBA := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: bounds.Dx(), Y: bounds.Dy()}})

	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			imgRGBA.Set(i, j, img.At(i, j))
		}
	}

	return imgRGBA
}

// GenerateMessage generates specified binary sequence.
func GenerateMessage(sequence string) []byte {
	message := make([]byte, len(sequence))

	for i, bit := range sequence {
		if bit == '1' {
			message[i] = 0x01
		} else {
			message[i] = 0x00
		}
	}

	return message
}

// GenerateRandomMessage generates random binary sequence with specified length.
func GenerateRandomMessage(length int) []byte {
	rMessage := make([]byte, length)

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	for i := 0; i < cap(rMessage); i++ {
		if r.Int()%2 == 1 {
			rMessage[i] = 0x00
		} else {
			rMessage[i] = 0x01
		}
	}

	return rMessage
}

// GenerateKey generates key of random intervals with specified length in the range from one to specified maximum
// interval step.
func GenerateKey(keyLen int, maxInterval int) []byte {
	key := make([]byte, keyLen)

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	for i := range key {
		key[i] = byte(r.Intn(maxInterval) + 1)
	}

	return key
}
