package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func GetDominant(img image.Image) color.Color {
	// build a map to rank
	cmap := make(map[color.Color]int)
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			cmap[img.At(x, y)]++
		}
	}
	// dominant color, score
	var dom color.Color
	score := 0
	// color, score
	for c, s := range cmap {
		if s > score {
			dom = c
			score = s
		}
	}
	return dom
}

func ColorStringfy(c color.Color) string {
	ri, gi, bi, _ := c.RGBA()
	r, g, b := int(ri)/256, int(gi)/256, int(bi)/256
	return fmt.Sprintf("rgb(%d,%d,%d)", r, b, g)
}

type Imagefile string

func (n Imagefile) Decode() image.Image {
	var img image.Image
	s := string(n)

	file, err := os.Open(s)
	if err != nil {
		fmt.Println("Failed to open file.")
		return nil
	}

	if strings.Contains(s, ".png") {
		img, _ = png.Decode(file)
	}
	if strings.Contains(s, ".jpg") || strings.Contains(s, ".jpeg") {
		img, _ = jpeg.Decode(file)
	}

	return img
}

func (n Imagefile) FindDominant() string {
	return ColorStringfy(GetDominant(n.Decode()))
}
func main() {
	if len(os.Args) > 1 {
		img_name := Imagefile(os.Args[1])
		fmt.Println(img_name.FindDominant())
	}
}
