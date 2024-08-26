package utilities

import (
	"image"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/chai2010/webp"

	_ "image/jpeg"
	_ "image/png"

	"github.com/google/uuid"
)

func init() {
	// Register the WebP decoder
	image.RegisterFormat("webp", "RIFF????WEBPVP8 ", webp.Decode, webp.DecodeConfig)
}

func SaveProfilePicture(file multipart.File) (string, error) {
	// Create a unique filename for the uploaded file
	fileName := uuid.New().String() + ".webp"
	filePath := filepath.Join("uploads", "images", "profile_pictures", fileName)

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// Make the image square
	squareImg := makeSquare(img)

	// Create the output file
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Encode the image as a WebP file with a quality of 75
	err = webp.Encode(outFile, squareImg, &webp.Options{Quality: 75})
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func makeSquare(img image.Image) image.Image {
	// Get image dimensions
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Calculate the size and position for the square crop
	var squareSize int
	var xOffset, yOffset int

	if width > height {
		squareSize = height
		xOffset = (width - height) / 2
		yOffset = 0
	} else {
		squareSize = width
		xOffset = 0
		yOffset = (height - width) / 2
	}

	// Crop the image to a square
	return img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(xOffset, yOffset, xOffset+squareSize, yOffset+squareSize))
}
