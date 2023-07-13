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
	Header        			*[]Header        			`json:"Header"`
	Item          			*[]Item          			`json:"Item"`
	ItemOperation 			*[]ItemOperation 			`json:"ItemOperation"`
	ItemOperationComponent	*[]ItemOperationComponent	`json:"ItemOperationComponent"`
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
	DestinationDeliverToParty                int     `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                string  `json:"DestinationDeliverToPlant"`
	DepartureDeliverFromParty                int     `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                string  `json:"DepartureDeliverFromPlant"`
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
	Product                                  string   `json:"Product"`
	Buyer                                    int      `json:"Buyer"`
	Seller                                   int      `json:"Seller"`
	DeliverToParty                           int      `json:"DeliverToParty"`
	DeliverToPlant                           string   `json:"DeliverToPlant"`
	DeliverFromParty                         int      `json:"DeliverFromParty"`
	DeliverFromPlant                         string   `json:"DeliverFromPlant"`
	ProductionPlantBusinessPartner           int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string   `json:"ProductionPlant"`
	OperationsText                           string   `json:"OperationsText"`
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
	CreationDate                             string   `json:"CreationDate"`
	LastChangeDate                           string   `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}

type ItemOperation struct {
	Operations                               int      `json:"Operations"`
	OperationsItem                           int      `json:"OperationsItem"`
	OperationID                              int      `json:"OperationID"`
	SupplyChainRelationshipID                int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int      `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string   `json:"Product"`
	Buyer                                    int      `json:"Buyer"`
	Seller                                   int      `json:"Seller"`
	DeliverToParty                           int      `json:"DeliverToParty"`
	DeliverToPlant                           string   `json:"DeliverToPlant"`
	DeliverFromParty                         int      `json:"DeliverFromParty"`
	DeliverFromPlant                         string   `json:"DeliverFromPlant"`
	ProductionPlantBusinessPartner           int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string   `json:"ProductionPlant"`
	Sequence                                 int      `json:"Sequence"`
	SequenceText                             *string  `json:"SequenceText"`
	OperationText                            string   `json:"OperationText"`
	OperationStatus                          *string  `json:"OperationStatus"`
	ResponsiblePlannerGroup                  *string  `json:"ResponsiblePlannerGroup"`
	OperationUnit                            string   `json:"OperationUnit"`
	StandardLotSizeQuantity                  *float32 `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity                   *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity                   *float32 `json:"MaximumLotSizeQuantity"`
	PlainLongText                            *string  `json:"PlainLongText"`
	WorkCenter                               *int     `json:"WorkCenter"`
	CapacityCategoryCode                     *string  `json:"CapacityCategoryCode"`
	OperationCostingRelevancyType            *string  `json:"OperationCostingRelevancyType"`
	OperationSetupType                       *string  `json:"OperationSetupType"`
	OperationSetupGroupCategory              *string  `json:"OperationSetupGroupCategory"`
	OperationSetupGroup                      *string  `json:"OperationSetupGroup"`
	OperationReferenceQuantity               *float32 `json:"OperationReferenceQuantity"`
	MaximumWaitDuration                      *float32 `json:"MaximumWaitDuration"`
	StandardWaitDuration                     *float32 `json:"StandardWaitDuration"`
	MinimumWaitDuration                      *float32 `json:"MinimumWaitDuration"`
	WaitDurationUnit                         *string  `json:"WaitDurationUnit"`
	MaximumQueueDuration                     *float32 `json:"MaximumQueueDuration"`
	StandardQueueDuration                    *float32 `json:"StandardQueueDuration"`
	MinimumQueueDuration                     *float32 `json:"MinimumQueueDuration"`
	QueueDurationUnit                        *string  `json:"QueueDurationUnit"`
	MaximumMoveDuration                      *float32 `json:"MaximumMoveDuration"`
	StandardMoveDuration                     *float32 `json:"StandardMoveDuration"`
	MinimumMoveDuration                      *float32 `json:"MinimumMoveDuration"`
	MoveDurationUnit                         *string  `json:"MoveDurationUnit"`
	StandardDeliveryDuration                 *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit             *string  `json:"StandardDeliveryDurationUnit"`
	StandardOperationScrapPercent            *float32 `json:"StandardOperationScrapPercent"`
	CostElement                              *string  `json:"CostElement"`
	ValidityStartDate                        *string  `json:"ValidityStartDate"`
	ValidityEndDate                          *string  `json:"ValidityEndDate"`
	CreationDate                             string   `json:"CreationDate"`
	LastChangeDate                           string   `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}

type ItemOperationComponent struct {
    Operations	                                    int	        `json:"Operations"`
    OperationsItem	                                int	        `json:"OperationsItem"`
    OperationID	                                    int	        `json:"OperationID"`
    BillOfMaterial	                                int	        `json:"BillOfMaterial"`
    BillOfMaterialItem	                            int	        `json:"BillOfMaterialItem"`
    SupplyChainRelationshipID	                    int         `json:"SupplyChainRelationshipID"`
    SupplyChainRelationshipDeliveryID	            int	        `json:"SupplyChainRelationshipDeliveryID"`
    SupplyChainRelationshipDeliveryPlantID	        int	        `json:"SupplyChainRelationshipDeliveryPlantID"`
    SupplyChainRelationshipStockConfPlantID	        int	        `json:"SupplyChainRelationshipStockConfPlantID"`
    ProductionPlantBusinessPartner	                int	        `json:"ProductionPlantBusinessPartner"`
    ProductionPlant	                                string	    `json:"ProductionPlant"`
    ComponentProduct	                            string	    `json:"ComponentProduct"`
    ComponentProductBuyer	                        int	        `json:"ComponentProductBuyer"`
    ComponentProductSeller	                        int	        `json:"ComponentProductSeller"`
    ComponentProductDeliverToParty	                int	        `json:"ComponentProductDeliverToParty"`
    ComponentProductDeliverToPlant	                string	    `json:"ComponentProductDeliverToPlant"`
    ComponentProductDeliverFromParty	            int	        `json:"ComponentProductDeliverFromParty"`
    ComponentProductDeliverFromPlant	            string	    `json:"ComponentProductDeliverFromPlant"`
    ComponentProductStandardQuantityInBaseUnit	    float32     `json:"ComponentProductStandardQuantityInBaseUnit"`
    ComponentProductStandardQuantityInDeliveryUnit	float32	    `json:"ComponentProductStandardQuantityInDeliveryUnit"`
    ComponentProductStandardScrapInPercent	        *float32	`json:"ComponentProductStandardScrapInPercent"`
    ComponentProductBaseUnit	                    string	    `json:"ComponentProductBaseUnit"`
    ComponentProductDeliveryUnit	                string	    `json:"ComponentProductDeliveryUnit"`
    StockConfirmationBusinessPartner	            int	        `json:"StockConfirmationBusinessPartner"`
    StockConfirmationPlant	                        string	    `json:"StockConfirmationPlant"`
    StockConfirmationPlantStorageLocation	        string	    `json:"StockConfirmationPlantStorageLocation"`
    IsMarkedForBackflush	                        *bool	    `json:"IsMarkedForBackflush"`
    ValidityStartDate	                            *string	    `json:"ValidityStartDate"`
    ValidityEndDate	                                *string	    `json:"ValidityEndDate"`
    CreationDate	                                string	    `json:"CreationDate"`
    LastChangeDate	                                string	    `json:"LastChangeDate"`
    IsMarkedForDeletion	                            *bool	    `json:"IsMarkedForDeletion"`
}
