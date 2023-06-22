package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header              *[]Header              `json:"Header"`
	Item                *[]Item                `json:"Item"`
	ComponentAllocation *[]ComponentAllocation `json:"ComponentAllocation"`
}

type Header struct {
	Operations                               int     `json:"Operations"`
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string  `json:"Product"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	DepartureDeliverFromParty                int     `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                string  `json:"DepartureDeliverFromPlant"`
	DestinationDeliverToParty                int     `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                string  `json:"DestinationDeliverToPlant"`
	OwnerProductionPlantBusinessPartner      int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     string  `json:"OwnerProductionPlant"`
	ProductBaseUnit                          string  `json:"ProductBaseUnit"`
	ProductDeliveryUnit                      string  `json:"ProductDeliveryUnit"`
	ProductProductionUnit                    string  `json:"ProductProductionUnit"`
	OperationsText                           string  `json:"OperationsText"`
	OperationsStatus                         *string `json:"OperationsStatus"`
	ResponsiblePlannerGroup                  *string `json:"ResponsiblePlannerGroup"`
	PlainLongText                            *string `json:"PlainLongText"`
	ValidityStartDate                        *string `json:"ValidityStartDate"`
	ValidityEndDate                          *string `json:"ValidityEndDate"`
	CreationDate                             string  `json:"CreationDate"`
	LastChangeDate                           string  `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
}

type Item struct {
	Operations                               int      `json:"Operations"`
	OperationsItem                           int      `json:"OperationsItem"`
	SupplyChainRelationshipID                int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int      `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  *string  `json:"Product"`
	Buyer                                    *int     `json:"Buyer"`
	Seller                                   *int     `json:"Seller"`
	DeliverFromParty                         *int     `json:"DeliverFromParty"`
	DeliverFromPlant                         *string  `json:"DeliverFromPlant"`
	DeliverToParty                           *int     `json:"DeliverToParty"`
	DeliverToPlant                           *string  `json:"DeliverToPlant"`
	ProductionPlantBusinessPartner           *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          *string  `json:"ProductionPlant"`
	OperationsText                           *string  `json:"OperationsText"`
	BillOfMaterial                           *int     `json:"BillOfMaterial"`
	OperationsStatus                         *string  `json:"OperationsStatus"`
	ResponsiblePlannerGroup                  *string  `json:"ResponsiblePlannerGroup"`
	OperationsUnit                           *string  `json:"OperationsUnit"`
	StandardLotSizeQuantity                  *float32 `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity                   *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity                   *float32 `json:"MaximumLotSizeQuantity"`
	PlainLongText                            *string  `json:"PlainLongText"`
	WorkCenter                               *int     `json:"WorkCenter"`
	ValidityStartDate                        *string  `json:"ValidityStartDate"`
	ValidityEndDate                          *string  `json:"ValidityEndDate"`
	CreationDate                             *string  `json:"CreationDate"`
	LastChangeDate                           *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}
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
