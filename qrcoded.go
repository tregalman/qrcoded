package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello QR Code")

	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = GenerateQRCode(file, "555-2368", Version(1))
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateQRCode generates a code based on the string passed in by the user
// and writes the QRCode version to a png file.
func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := 4*int(version) + 17
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	return png.Encode(w, img)
}

// Version is an alias of the int8 type
type Version int8
