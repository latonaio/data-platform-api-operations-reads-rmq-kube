package requests

type Header struct {
	Operations              int     `json:"Operations"`
	Product                 string  `json:"Product"`
	OwnerBusinessPartner    int     `json:"OwnerBusinessPartner"`
	OwnerPlant              string  `json:"OwnerPlant"`
	OperationsText          *string `json:"OperationsText"`
	OperationsStatus        *string `json:"OperationsStatus"`
	ResponsiblePlannerGroup *string `json:"ResponsiblePlannerGroup"`
	ValidityStartDate       *string `json:"ValidityStartDate"`
	ValidityEndDate         *string `json:"ValidityEndDate"`
	CreationDate            *string `json:"CreationDate"`
	LastChangeDate          *string `json:"LastChangeDate"`
	PlainLongText           *string `json:"PlainLongText"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}
