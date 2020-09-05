package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"
)

func loadImg(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func sqDiffUInt8(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}

func main() {
	fmt.Println(os.Getwd())
	img1, err := loadImg("banana.png")
	if err != nil {
		log.Fatal(err)
	}
	img2, err := loadImg("banana2.png")
	if err != nil {
		log.Fatal(err)
	}

	if img1, ok := img1.(*image.RGBA); ok {
		if img2, ok := img2.(*image.RGBA); ok {
			if img1.Bounds() != img2.Bounds() {
				fmt.Errorf("image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
			}

			accumError := int64(0)

			for i := 0; i < len(img1.Pix); i++ {
				accumError += int64(sqDiffUInt8(img1.Pix[i], img2.Pix[i]))
			}

			result := int64(math.Sqrt(float64(accumError)))

			fmt.Println(result)
		}
	}

}
