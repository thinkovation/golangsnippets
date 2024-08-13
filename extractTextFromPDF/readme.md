Certainly! An open-source alternative to `unipdf` for extracting text from PDF files in Go is the `rsc.io/pdf` library. This library is open source and can be used to extract text from PDFs.

First, you need to install the `rsc.io/pdf` library:

```sh
go get rsc.io/pdf
```

Then, you can write a function to extract text from a PDF file as shown below:

```go
package main

import (
	"fmt"
	"log"
	"strings"

	"rsc.io/pdf"
)

// ExtractTextFromPDF extracts text from a PDF file given its path.
func ExtractTextFromPDF(pdfPath string) (string, error) {
	// Open the PDF file
	doc, err := pdf.Open(pdfPath)
	if err != nil {
		return "", err
	}

	var text strings.Builder

	// Iterate over each page in the PDF
	for i := 0; i < doc.NumPage(); i++ {
		page := doc.Page(i)
		if page.V.IsNull() {
			continue
		}

		content := page.Content()
		for _, text := range content.Text {
			text.WriteString(text.S)
		}
	}

	return text.String(), nil
}

func main() {
	pdfPath := "path/to/your/file.pdf"
	text, err := ExtractTextFromPDF(pdfPath)
	if err != nil {
		log.Fatalf("Failed to extract text from PDF: %v", err)
	}

	fmt.Println("Extracted text:", text)
}
```

Here's a summary of the code:
1. Install the `rsc.io/pdf` library.
2. Define the `ExtractTextFromPDF` function to open a PDF file, read its pages, and extract the text from each page.
3. Open the PDF file and create a PDF reader.
4. Iterate over each page of the PDF and extract the text using the `pdf.Page` and `pdf.Content` methods.
5. Concatenate the extracted text from each page and return it.

To use this function, replace `"path/to/your/file.pdf"` with the actual path to your PDF file. The `main` function demonstrates how to call `ExtractTextFromPDF` and print the extracted text.