package scratchpad


import {
	"image"
	"os"
}

func main(csv) {
	// Open the file
	csvfile, err := os.Open(csv)
	if err != nil {
		log.Fatalln("Couldn not open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}
}

func ConvertImage(image)
	if image, ok := i.(*image.RGBA); ok {
		return i;
	}
}

func FastCompare(img1, img2 *image.RGBA) (int64, error) {
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