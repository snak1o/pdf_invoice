package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("", "mm", "A4", "")
	pdf.AddPage()

	//header

	pdf.SetFont("Times", "B", 20)
	//pdf.CellFormat(50, 10, "Invoice", "", 0, "BL", false, 0, "")
	pdf.Text(15, 15, "Invoice")

	pdf.SetFont("Times", "B", 15)
	//pdf.CellFormat(70, 10, "BoatCover.com", "", 0, "RB", false, 0, "http://localhost")
	pdf.Text(100, 15, "BoatCover.com")
	pdf.Ln(30)

	//body

	//left
	pdf.SetFont("Times", "", 12)
	pdf.Text(15, 35, "Firstname Lastname")
	pdf.Text(15, 39, "Street and number")
	pdf.Text(15, 43, "Zip City")
	pdf.Text(15, 47, "Country")
	pdf.Text(15, 51, "Email")
	pdf.Text(15, 55, "Phone number")

	//right
	pdf.Text(100, 35, "Date:")
	pdf.Text(100, 39, "Invoice number:")
	pdf.Text(100, 43, "Reference number:")
	pdf.Text(100, 47, "Bank account:")
	pdf.Text(100, 51, "BIC-code:")
	pdf.Text(100, 55, "Invoice due:")
	pdf.Text(100, 59, "Currency:")

	pdf.Text(140, 35, "25.09.2022")
	pdf.Text(140, 39, "0001")
	pdf.Text(140, 43, "43244")
	pdf.Text(140, 47, "FI43748234782374238")
	pdf.Text(140, 51, "NDEAFIHH")
	pdf.Text(140, 55, "30.09.2022")
	pdf.Text(140, 59, "EUR")

	//products
	pdf.SetFont("Times", "B", 10)
	pdf.SetLeftMargin(14)
	pdf.Ln(30)
	pdf.CellFormat(82, 10, "Product", "", 0, "L", false, 0, "")
	pdf.CellFormat(25, 10, "Quantity", "", 0, "R", false, 0, "")
	pdf.CellFormat(20, 10, "Price", "", 0, "R", false, 0, "")
	pdf.CellFormat(30, 10, "VAT (24%)", "", 0, "R", false, 0, "")
	pdf.CellFormat(25, 10, "Total price", "", 0, "R", false, 0, "")
	pdf.Line(15, 77, 195, 77)

	pdf.Ln(5)
	//item
	pdf.SetFont("Times", "", 12)
	pdf.CellFormat(82, 10, "Boat cover Buster 312 XXL PRO", "", 0, "L", false, 0, "")
	pdf.CellFormat(25, 10, "1", "", 0, "R", false, 0, "")
	pdf.CellFormat(20, 10, "80.65", "", 0, "R", false, 0, "")
	pdf.CellFormat(30, 10, "19.35", "", 0, "R", false, 0, "")
	pdf.CellFormat(25, 10, "100.00", "", 0, "R", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(82, 10, "Boat cover Buster 312 XXL PRO", "", 0, "L", false, 0, "")
	pdf.CellFormat(25, 10, "1", "", 0, "R", false, 0, "")
	pdf.CellFormat(20, 10, "80.65", "", 0, "R", false, 0, "")
	pdf.CellFormat(30, 10, "19.35", "", 0, "R", false, 0, "")
	pdf.CellFormat(25, 10, "100.00", "", 0, "R", false, 0, "")
	pdf.Ln(5)

	//total
	pdf.SetFont("Times", "", 13)
	pdf.Ln(20)
	pdf.CellFormat(157, 10, "Total amount without VAT:", "", 0, "R", false, 0, "")
	pdf.CellFormat(25, 10, "200.00", "", 0, "R", false, 0, "")
	pdf.Ln(6)
	pdf.CellFormat(157, 10, "VAT (24%):", "", 0, "R", false, 0, "")
	pdf.CellFormat(25, 10, "38.70", "", 0, "R", false, 0, "")
	pdf.Ln(6)
	pdf.CellFormat(157, 10, "Total amount:", "", 0, "R", false, 0, "")
	pdf.CellFormat(25, 10, "200.00", "", 0, "R", false, 0, "")

	//footer
	pdf.SetFont("Times", "", 10)
	pdf.Line(15, 270, 195, 270)
	pdf.Text(15, 274, "Boatcover Ltd.")
	pdf.Text(15, 278, "Address 13 box 31")
	pdf.Text(15, 282, "00432 Helsinki")
	pdf.Text(15, 286, "FINLAND")

	pdf.Text(75, 274, "Company ID: 4342342-6")
	pdf.Text(75, 278, "VAT-ID: FI4234234")
	pdf.Text(75, 282, "Phone: +358452614228")
	pdf.Text(75, 286, "E-mail: snakwcs@gmail.com")

	pdf.Text(160, 274, "Bank: Nordea")
	pdf.Text(160, 278, "FI43748234782374238")
	pdf.Text(160, 282, "NDEAFIHH")

	//create pdf
	err := pdf.OutputFileAndClose("invoice-number.pdf")
	if err != nil {
		log.Fatalf("Error saving pdf file: %s", err)
	}
}
