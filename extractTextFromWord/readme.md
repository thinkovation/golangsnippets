You are correct, I apologize for that oversight. Let's use the open-source library `github.com/baliance/gooxml` for extracting text from Word documents. 

Here’s how you can do it:

First, install the `gooxml` library:

```sh
go get github.com/baliance/gooxml
```

Then, you can write a function to extract text from a Word document as shown below:

```go
package main

import (
	"fmt"
	"log"

	"github.com/baliance/gooxml/document"
)

// ExtractTextFromWord extracts text from a Word document given its path.
func ExtractTextFromWord(docPath string) (string, error) {
	// Open the Word document
	doc, err := document.Open(docPath)
	if err != nil {
		return "", err
	}

	var text string

	// Iterate over each paragraph in the document
	for _, para := range doc.Paragraphs() {
		text += para.Text() + "\n"
	}

	return text, nil
}

func main() {
	docPath := "path/to/your/file.docx"
	text, err := ExtractTextFromWord(docPath)
	if err != nil {
		log.Fatalf("Failed to extract text from Word document: %v", err)
	}

	fmt.Println("Extracted text:", text)
}
```

Here's a summary of the code:
1. Install the `github.com/baliance/gooxml` library.
2. Define the `ExtractTextFromWord` function to open a Word document, read its paragraphs, and extract the text from each paragraph.
3. Open the Word document using `document.Open`.
4. Iterate over each paragraph in the document and concatenate the extracted text.
5. Return the extracted text.

To use this function, replace `"path/to/your/file.docx"` with the actual path to your Word document. The `main` function demonstrates how to call `ExtractTextFromWord` and print the extracted text.