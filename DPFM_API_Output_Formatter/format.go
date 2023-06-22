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
			&pm.DepartureDeliverFromParty,
			&pm.DepartureDeliverFromPlant,
			&pm.DestinationDeliverToParty,
			&pm.DestinationDeliverToPlant,
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
			DepartureDeliverFromParty:                data.DepartureDeliverFromParty,
			DepartureDeliverFromPlant:                data.DepartureDeliverFromPlant,
			DestinationDeliverToParty:                data.DestinationDeliverToParty,
			DestinationDeliverToPlant:                data.DestinationDeliverToPlant,
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
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
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
			DeliverFromParty:                         data.DeliverFromParty,
			DeliverFromPlant:                         data.DeliverFromPlant,
			DeliverToParty:                           data.DeliverToParty,
			DeliverToPlant:                           data.DeliverToPlant,
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

func ConvertToComponentAllocation(rows *sql.Rows) (*[]ComponentAllocation, error) {
	defer rows.Close()
	componentAllocation := make([]ComponentAllocation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ComponentAllocation{}

		err := rows.Scan(
			&pm.Operations,
			&pm.OperationsItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.IsMarkedForBackflush,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &componentAllocation, err
		}

		data := pm
		componentAllocation = append(componentAllocation, ComponentAllocation{
			Operations:           data.Operations,
			OperationsItem:       data.OperationsItem,
			BillOfMaterial:       data.BillOfMaterial,
			BillOfMaterialItem:   data.BillOfMaterialItem,
			IsMarkedForBackflush: data.IsMarkedForBackflush,
			ValidityStartDate:    data.ValidityStartDate,
			ValidityEndDate:      data.ValidityEndDate,
			CreationDate:         data.CreationDate,
			LastChangeDate:       data.LastChangeDate,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &componentAllocation, nil
	}

	return &componentAllocation, nil
}
