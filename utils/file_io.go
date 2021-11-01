package utils

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"io/ioutil"
	"os"
)

func ReadImage(fileName string) (*image.RGBA, error) {
	imgFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("image %q not found on disk: %v", fileName, err)
	}

	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	return rgba, nil
}

func ReadTextFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("file %q not found on disk: %v", fileName, err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return data, nil
}

func ReadJsonFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("file %q not found on disk: %v", fileName, err)
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	return byteValue, nil
}
