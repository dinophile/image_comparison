package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

func ParseCSV(file string) [][]string {
	imageStrings := [][]string{}

	csvFile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Could not open csv file", err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		imageStrings = append(imageStrings, record)
	}
	return imageStrings
}

func ImageCompare(img1, img2 *image.NRGBA) (int64, error) {
	if img1.Bounds() != img2.Bounds() {
		return 0, fmt.Errorf("image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
	}

	accumError := int64(0)

	for i := 0; i < len(img1.Pix); i++ {
		accumError += int64(sqDiffUInt8(img1.Pix[i], img2.Pix[i]))
	}

	return int64(math.Sqrt(float64(accumError))), nil

}

func sqDiffUInt8(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}

func main() {
	imagePairs := ParseCSV("input.csv")
	// remove csv headers
	imagePairs = imagePairs[1:]

	// fetch images with http, loop through each array within
	// imagePairs and pass each into image comparison algo
	// return value

	for _, imagePair := range imagePairs {
		for j, image := range imagePair {
			// retrieve image
			resp, err := http.Get(image)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()

			// save image temporarily (will remove later)
			filePath := "/tmp/image" + strconv.Itoa(j) + ".png"
			file, err := os.Create(filePath)
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()

			_, err = io.Copy(file, resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			// open image
			img1, err := os.Open("/tmp/image0.png")
			if err != nil {
				log.Fatalln(err)
			}
			defer img1.Close()

			// reset to beginning of memory block
			img1.Seek(0, 0)

			// get data from opened file and cast as type image.Image
			loadedImage1, err := png.Decode(img1)
			if err != nil {
				log.Fatalln(err)
			}

			img2, err := os.Open("/tmp/image1.png")
			if err != nil {
				log.Fatalln(err)
			}
			defer img2.Close()

			img2.Seek(0, 0)

			loadedImage2, err := png.Decode(img2)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%T\n", loadedImage1)
			fmt.Printf("%T\n", loadedImage2)

			// here I'm trying to do type coercion from image.Image to image.RGBA so I
			// can have access to the pixels and can pass to my compare algo
			if loadedImage1, ok := loadedImage1.(*image.NRGBA); ok {
				if loadedImage2, ok := loadedImage2.(*image.NRGBA); ok {
					ImageCompare(loadedImage1, loadedImage2)
				}
			}

			fmt.Printf("Hello")
		}
	}
}
