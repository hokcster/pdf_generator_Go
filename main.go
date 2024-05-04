package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	//pdf.Cell(40, 10, "Hello, World!")

	//pdf.ImageOptions("image.jpg", 10, 10, 30, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")
	//pdf.AddFont("Roboto", "", "Roboto-Regular.json")
	//pdf.SetFont("Roboto", "", 14)
	// Table headers
	headers := []string{"Country", "Capital", "Population"}

	// Table data
	data := [][]string{
		[]string{"China", "Beijing", "1,403M"},
		[]string{"India", "New Delhi", "1,366M"},
		[]string{"United States", "Washington, D.C.", "331M"},
	}

	pdf.SetFont("Arial", "B", 12)

	// Print table headers
	for _, header := range headers {
		pdf.CellFormat(40, 7, header, "1", 0, "", false, 0, "")
	}

	pdf.Ln(-1)

	// Print table data
	pdf.SetFont("Arial", "", 12)
	for _, line := range data {
		for _, cell := range line {
			pdf.CellFormat(40, 7, cell, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}

//pdf.ImageOptions("image.jpg", 10, 10, 30, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")
