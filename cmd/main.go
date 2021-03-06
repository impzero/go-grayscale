package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"

	"github.com/impzero/go-filters/filters"
)

func main() {
	// filtersMap is a map used to for filter input validation
	filtersMap := map[string]byte{
		"grayscale-coeff": 1,
		"grayscale-avg":   1,
		"bw":              1,
		"negative":        1,
	}

	img, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer img.Close()

	filename := path.Base(os.Args[1])
	filename = strings.TrimSuffix(filename, path.Ext(filename))

	imgType := path.Ext(os.Args[1])
	filter := os.Args[2]

	if _, ok := filtersMap[filter]; !ok {
		panic("go-filters: filter must be of type grayscale-coeff, grayscale-avg or bw")
	}

	var loadedImg image.Image
	if imgType == ".jpeg" || imgType == ".jpg" {
		loadedImg, err = jpeg.Decode(img)
	} else {
		loadedImg, err = png.Decode(img)
	}

	fmt.Println(filename)
	if err != nil {
		panic(err)
	}

	var filImg image.Image
	if strings.HasPrefix(filter, "grayscale") {
		filImg = filters.Grayscale(loadedImg, filter)
	} else if filter == "bw" {
		filImg = filters.BlackWhite(loadedImg)
	} else if filter == "negative" {
		filImg = filters.Negative(loadedImg)
	}

	filImgFile, err := os.Create(fmt.Sprintf("%s-%s%s", filename, filter, imgType))
	if err != nil {
		panic(err)
	}
	defer filImgFile.Close()

	err = jpeg.Encode(filImgFile, filImg, nil)

	if err != nil {
		panic(err)
	}
}
