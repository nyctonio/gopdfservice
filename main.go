package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jaytaylor/html2text"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	http.HandleFunc("/generate-pdf", htmlToPdfHandler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func htmlToPdfHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Close the request body after reading (important!)
	defer r.Body.Close()
	generatePdf(string(body))
	// modify to save to s3 and return url
	w.Write([]byte("PDF generated successfully."))
}

func generatePdf(htmlContent string) {
	// Convert HTML to plain text
	plainText, err := html2text.FromString(htmlContent)
	if err != nil {
		fmt.Println("Error converting HTML to plain text:", err)
		return
	}

	// Generate PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.MultiCell(0, 10, plainText, "", "", false)

	// Save the PDF to a file
	err = pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	fmt.Println("PDF generated successfully.")
}
