package requests

type ComponentAllocation struct {
	Operations           int     `json:"Operations"`
	OperationsItem       int     `json:"OperationsItem"`
	BillOfMaterial       int     `json:"BillOfMaterial"`
	BillOfMaterialItem   int     `json:"BillOfMaterialItem"`
	IsMarkedForBackflush *bool   `json:"IsMarkedForBackflush"`
	ValidityStartDate    *string `json:"ValidityStartDate"`
	ValidityEndDate      *string `json:"ValidityEndDate"`
	CreationDate         *string `json:"CreationDate"`
	LastChangeDate       *string `json:"LastChangeDate"`
}
