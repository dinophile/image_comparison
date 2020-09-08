package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
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

func ConvertImage(img image.Image) (*image.RGBA, error) {
	src := img
	bounds := img.Bounds()
	mask := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(mask, mask.Bounds(), src, bounds.Min, draw.Src)
	return mask, nil
}

func ImageCompare(img1, img2 image.Image) (int64, string, error) {
	start := time.Now()
	image1, err := ConvertImage(img1)
	if err != nil {
		log.Fatalln(err)
	}
	image2, err := ConvertImage(img2)
	if err != nil {
		log.Fatalln(err)
	}

	if image1.Bounds() != image2.Bounds() {
		return 0, fmt.Errorf("image bounds not equal: %+v, %+v", image1.Bounds(), image2.Bounds)
	}

	accumError := int64(0)

	for i := 0; i < len(image1.Pix); i++ {
		accumError += int64(sqDiffUInt8(image1.Pix[i], image2.Pix[i]))
	}

	elapsed := time.Since(start)

	return int64(math.Sqrt(float64(accumError))), elapsed, nil

}

func sqDiffUInt8(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}

func main() {
	// eventually this will take in a request body carrying a csv payload
	imagePairs := ParseCSV("input.csv")
	// remove csv headers
	imagePairs = imagePairs[1:]

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

			// do the same for image 2
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

			// pass to our comparison algorithm
			result, err := ImageCompare(loadedImage1, loadedImage2)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Print(result)

		}

		// every good script needs a sanity check! :D
		fmt.Printf("Hello")
	}
}
