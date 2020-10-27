package github.com/tolgaerdonmez/goutils

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/oliamb/cutter"
)

// ImageMetadata contains width and height of an image
type ImageMetadata struct {
	Width  int
	Height int
}

// TestImage extends image.Image
type TestImage struct {
	image.Image
	ImageMetadata
}

// Image interface for images
type Image interface {
	Save() func(string) error
	CropNumbers() func() (image.Image, error)
	CropAndSave() func(string) error
}

func getMetadata(i image.Image) ImageMetadata {
	var mdata ImageMetadata
	bounds := i.Bounds()
	mdata.Width = bounds.Dx()
	mdata.Height = bounds.Dy()
	return mdata
}

// GetImage returns your image and its metadata
func GetImage(filename string) (*TestImage, error) {
	mdata := ImageMetadata{}
	// Read image from file that already exists
	existingImageFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer existingImageFile.Close()

	loadedImage, err := jpeg.Decode(existingImageFile)
	if err != nil {
		return nil, err
	}
	mdata = getMetadata(loadedImage)
	img := &TestImage{loadedImage, mdata}
	return img, nil
}

func saveImage(filename string, i image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	jpeg.Encode(file, i, nil)
	return nil
}

// Save saves the image to given filename output
func (i TestImage) Save(filename string) error {
	return saveImage(filename, i)
}

// CropAndSave crops and saves to system using cutter
func (i *TestImage) CropAndSave(filename string, config cutter.Config) error {
	config.Options = cutter.Copy

	img, err := cutter.Crop(i, config)
	if err != nil {
		return err
	}
	return saveImage(filename, img)
}

// Crop crops and returns the new part
func (i *TestImage) Crop(config cutter.Config) (TestImage, error) {
	config.Options = cutter.Copy

	img, err := cutter.Crop(i, config)
	if err != nil {
		return TestImage{}, err
	}
	ni := TestImage{img, getMetadata(img)}
	return ni, nil
}

// RemoveBlank removes the blank part from the bottom of test image and saves to system by overriding original
func (i *TestImage) RemoveBlank(filename string) error {
	pixels, err := GetPixels(i)
	if err != nil {
		return err
	}

	firstI := -1
	notWhite := 0

	for i := len(pixels) - 1; i >= 0; i-- {
		row := pixels[i]
		if notWhite > len(pixels[0])*4/100 {
			// fmt.Println(firstI)
			break
		} else if firstI-i >= 10 {
			firstI = -1
		}
		for _, p := range row {
			if checkDarkness(p) {
				notWhite++
				if firstI == -1 {
					firstI = i
				}
			}
		}
	}

	config := cutter.Config{Width: i.Width, Height: firstI + 20}
	im, err := cutter.Crop(i, config)
	saveImage(filename, im)
	return err
}

// DetectStartingPixel detects the starting pixel from left to right
func (i *TestImage) DetectStartingPixel() int {
	pixels, err := GetPixels(i)
	if err != nil {
		return -1
	}

	start := -1
	for i := 0; i < len(pixels[0]); i++ {
		if start != -1 {
			break
		}
		for j := 0; j < len(pixels); j++ {
			p := pixels[j][i]
			if checkDarkness(p) {
				start = i
				break
			}
		}
	}
	return start
}
