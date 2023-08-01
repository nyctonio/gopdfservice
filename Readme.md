# HTML to PDF Converter

This is a simple Go program that provides a web server for converting HTML content to PDF and saving it to a file. The generated PDF file is saved with the name "output.pdf". The program uses the `jaytaylor/html2text` library to convert the HTML content to plain text before generating the PDF using the `jung-kurt/gofpdf` library.

## Prerequisites

To run this program, you need to have Go installed on your system. You can download and install Go from the official Go website: https://golang.org/

## Getting Started

- Clone the repository or copy the code to your local machine.
- Install the required dependencies by running the following command:

```bash
go get github.com/jaytaylor/html2text
go get github.com/jung-kurt/gofpdf
```

Run the program using the following command:

```go
go run main.go
```

The server will start listening on port 8080.

## Usage

Once the server is running, you can use a tool like `curl`, Postman, or any other HTTP client to send a POST request to the `/generate-pdf` endpoint with the HTML content in the request body. The server will convert the HTML content to plain text, generate a PDF, and save it to a file named "output.pdf" in the current directory.


```bash
curl -X POST -H "Content-Type: text/html" -d '<html><body><h1>Hello, World!</h1></body></html>' http://localhost:8080/generate-pdf
```

The program will print "PDF generated successfully." on the console, indicating that the PDF has been successfully created.

## Note

- The HTML content should be included in the request body as plain text.
- The server only accepts POST requests to the `/generate-pdf` endpoint.
- The generated PDF will be saved in the same directory as the executable with the filename "output.pdf".

Feel free to modify the program to save the generated PDF to a different location or integrate it with cloud storage services like Amazon S3 to store the PDFs remotely.