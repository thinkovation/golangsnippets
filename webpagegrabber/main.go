package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// Define command-line flags
	var url string
	var outputFile string
	var timeout int

	flag.StringVar(&url, "url", "", "URL of the dynamic page to fetch")
	flag.StringVar(&outputFile, "o", "out.html", "Output file to save the content")
	flag.IntVar(&timeout, "t", 60, "Timeout in seconds")

	flag.Parse()

	// Validate URL
	if url == "" {
		log.Fatal("URL must be provided")
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// Create a new browser context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Fetch dynamic content
	var content string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.OuterHTML("html", &content),
	)
	if err != nil {
		log.Fatalf("Failed to fetch content: %v", err)
	}

	// Save content to the output file
	err = os.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	fmt.Printf("Content successfully saved to %s\n", outputFile)
}
