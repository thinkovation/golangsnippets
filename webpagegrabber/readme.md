# Using chromedp
Chromedp is a Go library that provides a way to drive a headless Chrome/Chromium browser. Here's a basic example of how you might use chromedp to fetch the content of a dynamic page:

## 1. Install chromedp
You need to install the chromedp package. Run the following command:

```
go get -u github.com/chromedp/chromedp
```

## Running the code

### Command-Line Flags:

url: The URL of the dynamic page to fetch.
outputFile: The file to save the fetched content to, defaulting to out.html.
timeout: The timeout duration in seconds, defaulting to 60 seconds.

The dynamic content is fetched using chromedp and saved to the specified output file.

### Usage
To run the program, use the following command:
```
go run main.go -url "https://en.wikipedia.org/wiki/Blog" -o "output.html" -t 120
```

This command fetches content from https://en.wikipedia.org/wiki/Blog, saves it to output.html, and sets a timeout of 120 seconds.
If you omit the -o and -t flags, it will default to saving to out.html and using a 60-second timeout: