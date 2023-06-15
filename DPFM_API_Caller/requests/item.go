package requests

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
	DeliverFromParty                         int      `json:"DeliverFromParty"`
	DeliverFromPlant                         string   `json:"DeliverFromPlant"`
	DeliverToParty                           int      `json:"DeliverToParty"`
	DeliverToPlant                           string   `json:"DeliverToPlant"`
	ProductionPlantBusinessPartner           int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string   `json:"ProductionPlant"`
	OperationsText                           string   `json:"OperationsText"`
	BillOfMaterial                           *int     `json:"BillOfMaterial"`
	ValidityStartDate                        *string  `json:"ValidityStartDate"`
	ValidityEndDate                          *string  `json:"ValidityEndDate"`
	OperationsStatus                         *string  `json:"OperationsStatus"`
	ResponsiblePlannerGroup                  *string  `json:"ResponsiblePlannerGroup"`
	StandardLotSizeQuantity                  *float32 `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity                   *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity                   *float32 `json:"MaximumLotSizeQuantity"`
	OperationsUnit                           *string  `json:"OperationsUnit"`
	CreationDate                             *string  `json:"CreationDate"`
	LastChangeDate                           *string  `json:"LastChangeDate"`
	PlainLongText                            *string  `json:"PlainLongText"`
	WorkCenterTypeCode                       *string  `json:"WorkCenterTypeCode"`
	WorkCenterInternalID                     *string  `json:"WorkCenterInternalID"`
	CapacityCategoryCode                     *string  `json:"CapacityCategoryCode"`
	OperationCostingRelevancyType            *string  `json:"OperationCostingRelevancyType"`
	OperationScrapPercent                    *float32 `json:"OperationScrapPercent"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}
