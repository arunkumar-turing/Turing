package main

import (
	"fmt"
	"io/ioutil"

	"github.com/example/imageprocessor"
)

func main() {
	// Load image from file
	imgData, err := ioutil.ReadFile("example.jpg")
	if err != nil {
		panic(err)
	}

	// Process image asynchronously using functional callback
	result := imageprocessor.ProcessImageAsync(imgData,
		imageprocessor.Resize(500, 500),
		imageprocessor.ApplyFilter(imageprocessor.SepiaFilter),
		imageprocessor.SaveToFile("result.jpg"),
	)

	// Wait for the processing to complete and print the result
	fmt.Println(result)
}
