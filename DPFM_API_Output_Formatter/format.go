package dpfm_api_output_formatter

import (
	"data-platform-api-operations-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		pm := &requests.Header{}
		i++

		err := rows.Scan(
			&pm.Operations,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.DestinationDeliverToParty,
			&pm.DestinationDeliverToPlant,
			&pm.DepartureDeliverFromParty,
			&pm.DepartureDeliverFromPlant,
			&pm.OwnerProductionPlantBusinessPartner,
			&pm.OwnerProductionPlant,
			&pm.ProductBaseUnit,
			&pm.ProductDeliveryUnit,
			&pm.ProductProductionUnit,
			&pm.OperationsText,
			&pm.OperationsStatus,
			&pm.ResponsiblePlannerGroup,
			&pm.PlainLongText,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
			Operations:                               data.Operations,
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:        data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:   data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Product:                                  data.Product,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			DestinationDeliverToParty:                data.DestinationDeliverToParty,
			DestinationDeliverToPlant:                data.DestinationDeliverToPlant,
			DepartureDeliverFromParty:                data.DepartureDeliverFromParty,
			DepartureDeliverFromPlant:                data.DepartureDeliverFromPlant,
			OwnerProductionPlantBusinessPartner:      data.OwnerProductionPlantBusinessPartner,
			OwnerProductionPlant:                     data.OwnerProductionPlant,
			ProductBaseUnit:                          data.ProductBaseUnit,
			ProductDeliveryUnit:                      data.ProductDeliveryUnit,
			ProductProductionUnit:                    data.ProductProductionUnit,
			OperationsText:                           data.OperationsText,
			OperationsStatus:                         data.OperationsStatus,
			ResponsiblePlannerGroup:                  data.ResponsiblePlannerGroup,
			PlainLongText:                            data.PlainLongText,
			ValidityStartDate:                        data.ValidityStartDate,
			ValidityEndDate:                          data.ValidityEndDate,
			CreationDate:                             data.CreationDate,
			LastChangeDate:                           data.LastChangeDate,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.Operations,
			&pm.OperationsItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.OperationsText,
			&pm.BillOfMaterial,
			&pm.OperationsStatus,
			&pm.ResponsiblePlannerGroup,
			&pm.OperationsUnit,
			&pm.StandardLotSizeQuantity,
			&pm.MinimumLotSizeQuantity,
			&pm.MaximumLotSizeQuantity,
			&pm.PlainLongText,
			&pm.WorkCenter,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}

		data := pm
		item = append(item, Item{
			Operations:                               data.Operations,
			OperationsItem:                           data.OperationsItem,
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:        data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:   data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Product:                                  data.Product,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			DeliverToParty:                           data.DeliverToParty,
			DeliverToPlant:                           data.DeliverToPlant,
			DeliverFromParty:                         data.DeliverFromParty,
			DeliverFromPlant:                         data.DeliverFromPlant,
			ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
			ProductionPlant:                          data.ProductionPlant,
			OperationsText:                           data.OperationsText,
			BillOfMaterial:                           data.BillOfMaterial,
			OperationsStatus:                         data.OperationsStatus,
			ResponsiblePlannerGroup:                  data.ResponsiblePlannerGroup,
			OperationsUnit:                           data.OperationsUnit,
			StandardLotSizeQuantity:                  data.StandardLotSizeQuantity,
			MinimumLotSizeQuantity:                   data.MinimumLotSizeQuantity,
			MaximumLotSizeQuantity:                   data.MaximumLotSizeQuantity,
			PlainLongText:                            data.PlainLongText,
			WorkCenter:                               data.WorkCenter,
			ValidityStartDate:                        data.ValidityStartDate,
			ValidityEndDate:                          data.ValidityEndDate,
			CreationDate:                             data.CreationDate,
			LastChangeDate:                           data.LastChangeDate,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}

func ConvertToItemOperation(rows *sql.Rows) (*[]ItemOperation, error) {
	defer rows.Close()
	itemOperation := make([]ItemOperation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemOperation{}

		err := rows.Scan(
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.Sequence,
			&pm.SequenceText,
			&pm.OperationText,
			&pm.OperationStatus,
			&pm.ResponsiblePlannerGroup,
			&pm.OperationUnit,
			&pm.StandardLotSizeQuantity,
			&pm.MinimumLotSizeQuantity,
			&pm.MaximumLotSizeQuantity,
			&pm.PlainLongText,
			&pm.WorkCenter,
			&pm.CapacityCategoryCode,
			&pm.OperationCostingRelevancyType,
			&pm.OperationSetupType,
			&pm.OperationSetupGroupCategory,
			&pm.OperationSetupGroup,
			&pm.OperationReferenceQuantity,
			&pm.MaximumWaitDuration,
			&pm.StandardWaitDuration,
			&pm.MinimumWaitDuration,
			&pm.WaitDurationUnit,
			&pm.MaximumQueueDuration,
			&pm.StandardQueueDuration,
			&pm.MinimumQueueDuration,
			&pm.QueueDurationUnit,
			&pm.MaximumMoveDuration,
			&pm.StandardMoveDuration,
			&pm.MinimumMoveDuration,
			&pm.MoveDurationUnit,
			&pm.StandardDeliveryDuration,
			&pm.StandardDeliveryDurationUnit,
			&pm.StandardOperationScrapPercent,
			&pm.CostElement,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemOperation, err
		}

		data := pm
		itemOperation = append(itemOperation, ItemOperation{

			Operations:                               data.Operations,
			OperationsItem:                           data.OperationsItem,
			OperationID:                              data.OperationID,
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:        data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:   data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Product:                                  data.Product,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			DeliverToParty:                           data.DeliverToParty,
			DeliverToPlant:                           data.DeliverToPlant,
			DeliverFromParty:                         data.DeliverFromParty,
			DeliverFromPlant:                         data.DeliverFromPlant,
			ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
			ProductionPlant:                          data.ProductionPlant,
			Sequence:                                 data.Sequence,
			SequenceText:                             data.SequenceText,
			OperationText:                            data.OperationText,
			OperationStatus:                          data.OperationStatus,
			ResponsiblePlannerGroup:                  data.ResponsiblePlannerGroup,
			OperationUnit:                            data.OperationUnit,
			StandardLotSizeQuantity:                  data.StandardLotSizeQuantity,
			MinimumLotSizeQuantity:                   data.MinimumLotSizeQuantity,
			MaximumLotSizeQuantity:                   data.MaximumLotSizeQuantity,
			PlainLongText:                            data.PlainLongText,
			WorkCenter:                               data.WorkCenter,
			CapacityCategoryCode:                     data.CapacityCategoryCode,
			OperationCostingRelevancyType:            data.OperationCostingRelevancyType,
			OperationSetupType:                       data.OperationSetupType,
			OperationSetupGroupCategory:              data.OperationSetupGroupCategory,
			OperationSetupGroup:                      data.OperationSetupGroup,
			OperationReferenceQuantity:               data.OperationReferenceQuantity,
			MaximumWaitDuration:                      data.MaximumWaitDuration,
			StandardWaitDuration:                     data.StandardWaitDuration,
			MinimumWaitDuration:                      data.MinimumWaitDuration,
			WaitDurationUnit:                         data.WaitDurationUnit,
			MaximumQueueDuration:                     data.MaximumQueueDuration,
			StandardQueueDuration:                    data.StandardQueueDuration,
			MinimumQueueDuration:                     data.MinimumQueueDuration,
			QueueDurationUnit:                        data.QueueDurationUnit,
			MaximumMoveDuration:                      data.MaximumMoveDuration,
			StandardMoveDuration:                     data.StandardMoveDuration,
			MinimumMoveDuration:                      data.MinimumMoveDuration,
			MoveDurationUnit:                         data.MoveDurationUnit,
			StandardDeliveryDuration:                 data.StandardDeliveryDuration,
			StandardDeliveryDurationUnit:             data.StandardDeliveryDurationUnit,
			StandardOperationScrapPercent:            data.StandardOperationScrapPercent,
			CostElement:                              data.CostElement,
			ValidityStartDate:                        data.ValidityStartDate,
			ValidityEndDate:                          data.ValidityEndDate,
			CreationDate:                             data.CreationDate,
			LastChangeDate:                           data.LastChangeDate,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemOperation, nil
	}

	return &itemOperation, nil
}

func ConvertToItemOperationComponent(rows *sql.Rows) (*[]ItemOperationComponent, error) {
	defer rows.Close()
	itemOperationComponent := make([]ItemOperationComponent, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemOperationComponent{}

		err := rows.Scan(
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ComponentProduct,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ComponentProductDeliverToParty,
			&pm.ComponentProductDeliverToPlant,
			&pm.ComponentProductDeliverFromParty,
			&pm.ComponentProductDeliverFromPlant,
			&pm.ComponentProductStandardQuantityInBaseUnit,
			&pm.ComponentProductStandardQuantityInDeliveryUnit,
			&pm.ComponentProductStandardScrapInPercent,
			&pm.ComponentProductBaseUnit,
			&pm.ComponentProductDeliveryUnit,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantStorageLocation,
			&pm.IsMarkedForBackflush,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemOperationComponent, err
		}

		data := pm
		itemOperationComponent = append(itemOperationComponent, ItemOperationComponent{

			Operations:										data.Operations,
			OperationsItem:									data.OperationsItem,
			OperationID:									data.OperationID,
			BillOfMaterial:									data.BillOfMaterial,
			BillOfMaterialItem:								data.BillOfMaterialItem,
			SupplyChainRelationshipID:						data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:				data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:			data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:		data.SupplyChainRelationshipStockConfPlantID,
			ProductionPlantBusinessPartner:					data.ProductionPlantBusinessPartner,
			ProductionPlant:								data.ProductionPlant,
			ComponentProduct:								data.ComponentProduct,
			ComponentProductBuyer:							data.ComponentProductBuyer,
			ComponentProductSeller:							data.ComponentProductSeller,
			ComponentProductDeliverToParty:					data.ComponentProductDeliverToParty,
			ComponentProductDeliverToPlant:					data.ComponentProductDeliverToPlant,
			ComponentProductDeliverFromParty:				data.ComponentProductDeliverFromParty,
			ComponentProductDeliverFromPlant:				data.ComponentProductDeliverFromPlant,
			ComponentProductStandardQuantityInBaseUnit:		data.ComponentProductStandardQuantityInBaseUnit,
			ComponentProductStandardQuantityInDeliveryUnit:	data.ComponentProductStandardQuantityInDeliveryUnit,
			ComponentProductStandardScrapInPercent:			data.ComponentProductStandardScrapInPercent,
			ComponentProductBaseUnit:						data.ComponentProductBaseUnit,
			ComponentProductDeliveryUnit:					data.ComponentProductDeliveryUnit,
			StockConfirmationBusinessPartner:				data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:							data.StockConfirmationPlant,
			StockConfirmationPlantStorageLocation:			data.StockConfirmationPlantStorageLocation,
			IsMarkedForBackflush:							data.IsMarkedForBackflush,
			ValidityStartDate:								data.ValidityStartDate,
			ValidityEndDate:								data.ValidityEndDate,
			CreationDate:									data.CreationDate,
			LastChangeDate:									data.LastChangeDate,
			IsMarkedForDeletion:							data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemOperationComponent, nil
	}

	return &itemOperationComponent, nil
}
