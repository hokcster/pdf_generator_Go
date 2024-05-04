package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, World!")

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
	pdf.ImageOptions("image.jpg", 10, 10, 30, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")
	pdf.AddFont("Roboto", "", "Roboto-Regular.json")
	pdf.SetFont("Roboto", "", 14)
}

//pdf.ImageOptions("image.jpg", 10, 10, 30, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")
