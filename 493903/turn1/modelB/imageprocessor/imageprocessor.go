package imageprocessor

import (
	"image"
	"image/color"
	"image/draw"
	"os"
	"path/filepath"

	"golang.org/x/image/math/f64"
)

// Callback represents a function that takes an image and returns a new image.
type Callback func(image.Image) image.Image

// ProcessImageAsync processes an image using a chain of asynchronous callbacks.
func ProcessImageAsync(imgData []byte, callbacks ...Callback) error {
	// Decode the image
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return err
	}

	// Create a channel to receive results from each callback
	resultChan := make(chan image.Image, len(callbacks))

	// Apply each callback asynchronously using goroutines
	for _, callback := range callbacks {
		go func(cb Callback) {
			resultChan <- cb(img)
		}(callback)
	}

	// Wait for all results to be received
	for range callbacks {
		img = <-resultChan
	}
