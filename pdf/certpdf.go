package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"

	"training.go/GenCert/cert"
)

type PdfSaver struct {
	OutputDir string
}

//New prend en parametre le chemin du dossier,
//renvoie un pointeur PdfSaver et une erreur
func New(outputDir string) (*PdfSaver, error) {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		var p *PdfSaver
		return p, err
	}
	p := &PdfSaver{
		OutputDir: outputDir,
	}

	return p, nil
}

//Save prend en parameter un certificat,
//renvoie une erreur,
//est liée a l'interface Saver
func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// Background
	background(pdf)

	// -----------------
	// Header
	header(pdf, &cert)
	pdf.Ln(30)

	// ------------
	// Body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	// Body - Student Name
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	// Body - Participation
	pdf.SetFont("Times", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	// Body - Date
	pdf.SetFont("Times", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	// footer
	footer(pdf)

	// save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Certificat sauvegarder à '%v'\n", path)
	return nil

}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions("img/background.png",
		0, 0,
		pageWidth, pageHeight,
		false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	margin := 25.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/gopher.png"
	pdf.ImageOptions(filename,
		x+margin, 20,
		imageWidth, 0,
		false, opts, 0, "")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename,
		x-margin, 20,
		imageWidth, 0,
		false, opts, 0, "")
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	pageWidth, pageHeight := pdf.GetPageSize()
	imageWidth := 50.0
	x := pageWidth - imageWidth - 20.0
	y := pageHeight - imageWidth - 10.0

	pdf.ImageOptions("img/stamp.png", x, y, imageWidth, 0, false, opts, 0, "")
}
