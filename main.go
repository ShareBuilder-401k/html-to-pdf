package main

import (
	"encoding/json"
	"fmt"
	"html-to-pdf/model"
	"html-to-pdf/pdf"
	"html-to-pdf/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Context struct {
	WorkingDir  string
	Templates   template.HTMLTemplateInterface
	PDF         pdf.PDFProviderInterface
	Development bool
}

func (c *Context) handleRequest(w http.ResponseWriter, r *http.Request) {
	j := json.NewDecoder(r.Body)
	pdfModel := model.NewPDFModel(c.WorkingDir)
	err := j.Decode(pdfModel)
	pdfModel.FillMissing()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformed request: " + err.Error()))
		return
	}

	if c.Development {
		// Refresh the Template Each Test
		c.Templates = template.NewHTMLTemplate()
	}

	t, err := c.Templates.Compile(pdfModel, "PDF")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating purchase agreement template: " + err.Error()))
		return
	}

	pdf, err := c.PDF.CreatePDF(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating purchase agreement PDF: " + err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdf)
}

func main() {
	wd, _ := os.Getwd()
	t := template.NewHTMLTemplate()
	p := pdf.NewPDFProvider()
	c := &Context{
		PDF:         p,
		Templates:   t,
		WorkingDir:  wd,
		Development: true,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Lets create some PDFs!\n")
	})

	r.HandleFunc("/pdf", c.handleRequest).Methods("POST")

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
