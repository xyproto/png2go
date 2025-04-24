package main

import (
	"errors"
	"image"
	"log"

	"github.com/xyproto/carveimg"
)

// Load the given filename, return an image, the width, height and an error if there are issues
func Load(filename string) (*image.NRGBA, int, int, error) {
	img, err := carveimg.LoadImage(filename)
	if err != nil {
		return nil, 0, 0, err
	}
	imageHeight := float64(img.Bounds().Max.Y - img.Bounds().Min.Y)
	if imageHeight == 0 {
		return nil, 0, 0, errors.New("the height of the given image is 0")
	}
	imageWidth := float64(img.Bounds().Max.X - img.Bounds().Min.X)
	if imageWidth == 0 {
		return nil, 0, 0, errors.New("the width of the given image is 0")
	}
	return img, int(imageWidth), int(imageHeight), nil
}

func main() {
	// 1. Read in an image
	img, w, h, err := Load("board.png")
	if err != nil {
		log.Fatalln(err)
	}

	_ = w
	_ = h
	_ = img

	// 2. Find the X positions of all the lines
	// 3. Find the Y positions of all the lines
	// 4. Find a formula for finding the lines between the first and last X line (linear? increasing in width?)
	// 5. Find a formula for finding the lines between the first and last Y line
	// 6. Find the size of the board, and check if there are intersections at the expected locations
	// 7. Find all the black stones by looking at the intersections
	// 8. Find all the white stones by looking at the intersections
}
