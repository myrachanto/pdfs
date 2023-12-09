package src

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

var businesPartner = BusinessPartner{SerialNo: "PD100", FullName: "P KIMANI", Gender: "Male", MaritalStatus: "Married", Dob: "1972", DocType: "National ID", MonthlyIncome: 300000, BusinessName: "PEWA ENTERPRISES"}
var loan = LoanApplicationRequest{
	LoanApplication: LoanApplication{CustomerId: 1, ProductId: 18, InterestRate: 25, RepaymentPeriod: 6, AmountApplied: 600000, Status: 0, DocDate: time.Now(), CustodyExemption: "N"},
	Collaterals: []Collateral{
		{ApplicationId: 55, CustomerId: 1, RefNo: "200081200963", AssetDesc: "KfN 366Z", AssetCategory: "Log Book", AssetLocation: "MILIMANI OPPOSITE", AppraisalValue: 200000, AssetId: "n/a"},
		{ApplicationId: 55, CustomerId: 1, RefNo: "300081200963", AssetDesc: "KCD 366Z", AssetCategory: "Log Book", AppraisalValue: 1000000, AssetId: "n/a"},
	},
}

func GeneratePdf(file2 string) error {
	fmt.Println("-------------------------------step 2")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetHeaderFunc(func() {
		// Set font for the header
		pdf.SetFont("Arial", "B", 12)

		// Add content to the header
		pdf.CellFormat(190, 10, "Vamos", "0", 0, "C", false, 0, "")
	})
	pdf.ImageOptions(
		"public/logo.jpg",
		10, 20,
		80, 50,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: false},
		0,
		"",
	)
	pdf.SetFooterFunc(func() {
		// Set font for the footer
		pdf.SetFont("Arial", "I", 8)
		// Set position for the footer at the bottom
		pdf.SetY(-25)

		// Add content to the footer
		pdf.CellFormat(190, 10, "This is a computer generated Invoice and as such does not have a signature.", "0", 0, "C", false, 0, "")
		pdf.Ln(5)
		pdf.CellFormat(190, 10, "Stalis pos", "0", 0, "C", false, 0, "")
	})
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(211, 211, 211)
	pdf.CellFormat(190, 7, "Stalis Pos", "0", 0, "CM", false, 0, "")
	// pdf.Image("public/imgs/business/acer.jpg",20, 30,40, 40,false, "", 0, "")
	pdf.Ln(10)
	//header one
	fmt.Println("-------------------------------step 4")
	pdf.SetXY(100, 20)
	pdf.MultiCell(90, 15, "Invoice", "0", "", true)
	pdf.SetXY(100, 35)
	pdf.MultiCell(90, 7, businesPartner.FullName, "0", "", true)

	pdf.Ln(12)
	pdf.SetXY(100, 45)
	pdf.MultiCell(90, 7, businesPartner.BusinessName, "0", "", true)
	pdf.Ln(10)
	pdf.SetXY(100, 55)
	pdf.CellFormat(50, 20, businesPartner.BusinessLocation, "0", 0, "TL", false, 0, "")
	pdf.Ln(10)
	pdf.SetXY(100, 60)
	pdf.CellFormat(50, 20, businesPartner.EmployerPhone, "0", 0, "TL", false, 0, "")
	pdf.Ln(10)
	pdf.SetXY(100, 65)
	pdf.CellFormat(50, 20, businesPartner.Email, "0", 0, "TL", false, 0, "")
	pdf.Ln(5)
	//header two
	pdf.Ln(20)
	pdf.SetXY(10, 70)
	pdf.MultiCell(90, 7, fmt.Sprintf("Amount Applied %d", loan.LoanApplication.AmountApplied), "0", "", false)
	pdf.Ln(10)
	// pdf.SetXY(10, 80)
	// pdf.CellFormat(50, 10, invoice.Location+", "+invoice.Street+", "+invoice.House, "0", 0, "TL", false, 0, "")
	// pdf.Ln(10)
	pdf.SetXY(10, 85)
	pdf.CellFormat(50, 10, fmt.Sprintf("Repayment Period %d", loan.LoanApplication.RepaymentPeriod), "0", 0, "TL", false, 0, "")
	pdf.Ln(10)
	pdf.SetXY(10, 90)
	pdf.CellFormat(50, 10, fmt.Sprintf("Interest Rate  %d", loan.LoanApplication.InterestRate), "0", 0, "TL", false, 0, "")
	pdf.Ln(10)
	pdf.CellFormat(10, 20, "Collaterals:    ", "0", 0, "TL", false, 0, "")
	pdf.Ln(5)
	pdf.MultiCell(180, 7, "Asset Name                                                                                                        Asset Appraisal                                      Asset location", "0", "", true)
	pdf.SetFont("Arial", "", 8)
	for _, v := range loan.Collaterals {
		pdf.Ln(5)
		pdf.CellFormat(60, 20, v.AssetDesc, "0", 0, "TL", false, 0, "")
		pdf.CellFormat(60, 20, fmt.Sprintln(v.AppraisalValue), "0", 0, "TR", false, 0, "")
		// pdf.CellFormat(20, 20, v.Landmark, "0", 0, "TC", false, 0, "")
		pdf.CellFormat(80, 20, v.AssetLocation, "0", 0, "TC", false, 0, "")
	}
	pdf.SetFont("Arial", "B", 8)
	// pdf.Ln(10)
	// pdf.CellFormat(150, 20, "Subtotal", "0", 0, "TR", false, 0, "")
	// pdf.CellFormat(30, 20, business.Currency+fmt.Sprintln(invoice.Subtotal)+".00", "0", 0, "TR", false, 0, "")
	// pdf.Ln(8)
	// pdf.MultiCell(150, 7, fmt.Sprintln(loan.LoanApplication.AmountApplied), "0", "", true)
	// pdf.CellFormat(130, 20, "Total", fmt.Sprintln(invoice.Amount), 0, "TR", false, 0, "")
	// pdf.CellFormat(30, 20, business.Currency+fmt.Sprintln(invoice.Total)+".00", "0", 0, "TR", false, 0, "")
	pdf.Ln(10)
	pdf.CellFormat(10, 20, "Invoice Delivery Date:    ", "0", 0, "TL", false, 0, "")
	pdf.Ln(5)
	pdf.MultiCell(180, 7, fmt.Sprintln(loan.LoanApplication.DocDate), "0", "", true)
	pdf.Ln(10)

	pdf.Ln(10)
	pdf.CellFormat(10, 20, "Instructions:    ", "0", 0, "TL", false, 0, "")
	pdf.Ln(5)
	// pdf.MultiCell(180, 7, invoice.Instructions, "0", "", true)
	// pdf.Ln(10)
	// Set auto page break to true with a large bottom margin
	// pdf.SetAutoPageBreak(true, 70)
	// fmt.Println("-------------------------------step 5", file2)
	return pdf.OutputFileAndClose(file2)

}
