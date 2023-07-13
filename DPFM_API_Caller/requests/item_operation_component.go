package requests

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
