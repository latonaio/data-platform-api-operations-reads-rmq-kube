package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Header            Header   `json:"Operations"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Header struct {
	Operations                               int     `json:"Operations"`
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  *string `json:"Product"`
	Buyer                                    *int    `json:"Buyer"`
	Seller                                   *int    `json:"Seller"`
	DepartureDeliverFromParty                *int    `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                *string `json:"DepartureDeliverFromPlant"`
	DestinationDeliverToParty                *int    `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                *string `json:"DestinationDeliverToPlant"`
	OwnerProductionBusinessPartner           *int    `json:"OwnerProductionBusinessPartner"`
	OwnerProductionPlant                     *string `json:"OwnerProductionPlant"`
	OperationsText                           *string `json:"OperationsText"`
	OperationsStatus                         *string `json:"OperationsStatus"`
	ResponsiblePlannerGroup                  *string `json:"ResponsiblePlannerGroup"`
	PlainLongText                            *string `json:"PlainLongText"`
	ValidityStartDate                        *string `json:"ValidityStartDate"`
	ValidityEndDate                          *string `json:"ValidityEndDate"`
	CreationDate                             *string `json:"CreationDate"`
	LastChangeDate                           *string `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
	Item                                     []Item  `json:"Item"`
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
