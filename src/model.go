package src

import (
	"time"

	"gorm.io/gorm"
)

type BusinessPartner struct {
	gorm.Model
	ID                 int     `json:"id"`
	SerialNo           string  `json:"serialNo"`
	FullName           string  `json:"fullName"`
	MaritalStatus      string  `json:"maritalStatus"`
	Dob                string  `json:"dob"`
	PinNo              string  `json:"pinNo"`
	CreditLimit        int     `json:"creditLimit"`
	DocType            string  `json:"docType"`
	DocNo              string  `json:"docNo"`
	PhoneNo            string  `json:"phoneNo"`
	Gender             string  `json:"gender"`
	PostalAddress      string  `json:"postalAddress"`
	PhysicalAddress    string  `json:"physicalAddress"`
	Email              string  `json:"email"`
	Landmark           string  `json:"landmark"`
	NokName            string  `json:"nokName"`
	NokDocType         string  `json:"nokDocType"`
	NokDocNo           string  `json:"nokDocNo"`
	NokPhoneNo         string  `json:"nokPhoneNo"`
	Relationship       string  `json:"relationship"`
	NokPhysicalAddress string  `json:"nokPhysicalAddress"`
	IncomeSource       string  `json:"incomeSource"`
	Designation        string  `json:"designation"`
	EmployeeNo         string  `json:"employeeNo"`
	EmployerName       string  `json:"employerName"`
	EmployerPhone      string  `json:"employerPhone"`
	OfficeLocation     string  `json:"officeLocation"`
	BusinessLocation   string  `json:"businessLocation"`
	BusinessNature     string  `json:"businessNature"`
	MonthlyIncome      float64 `json:"monthlyIncome"`
	BusinessName       string  `json:"businessName"`
	Overpayment        float64 `json:"overpayment"`
}

type LoanApplication struct {
	gorm.Model
	ID         int             `json:"id"`
	CustomerId int             `json:"customerId"`
	Customer   BusinessPartner `json:"customer" gorm:"foreignKey:CustomerId"`
	ProductId  int             `json:"productId"`
	// Product          Account         `json:"product" gorm:"foreignKey:ProductId"`
	InterestRate     int          `json:"interestRate"`
	RepaymentPeriod  int          `json:"repaymentPeriod"`
	AmountApplied    int          `json:"amountApplied"`
	Status           int          `json:"status"`
	DocDate          time.Time    `json:"DocDate"`
	CustodyExemption string       `json:"custodyExemption"`
	Collateral       []Collateral `json:"Collateral" gorm:"foreignKey:ApplicationId"`
	// LoanDisbursement []LoanDisbursement `json:"LoanDisbursement" gorm:"foreignKey:ApplicationId"`
}
type LoanApplicationRequest struct {
	LoanApplication LoanApplication `json:"loanDetails"`
	Collaterals     []Collateral    `json:"collaterals"`
}
type LoanApplicationInfo struct {
	LoanApplication  LoanApplication `json:"loanDetails"`
	Collaterals      []Collateral    `json:"collaterals"`
	LoanDisbursement float64         `json:"totalDisbursements"`
}
type Collateral struct {
	gorm.Model
	ApplicationId  int    `json:"applicationId"`
	AssetId        string `gorm:"-" json:"assetId"`
	CustomerId     int    `json:"customerId"`
	RefNo          string `json:"refNo"`
	AssetDesc      string `json:"assetDesc"`
	AssetCategory  string `json:"assetCategory"`
	AssetLocation  string `json:"assetLocation"`
	IsCharged      string `json:"isCharged"`
	Chargee        string `json:"chargee"`
	Landmark       string `json:"landmark"`
	AppraisalValue int    `json:"appraisalValue"`
}
