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
			&pm.Product,
			&pm.OwnerBusinessPartner,
			&pm.OwnerPlant,
			&pm.OperationsText,
			&pm.OperationsStatus,
			&pm.ResponsiblePlannerGroup,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.PlainLongText,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
			Operations:              data.Operations,
			Product:                 data.Product,
			OwnerBusinessPartner:    data.OwnerBusinessPartner,
			OwnerPlant:              data.OwnerPlant,
			OperationsText:          data.OperationsText,
			OperationsStatus:        data.OperationsStatus,
			ResponsiblePlannerGroup: data.ResponsiblePlannerGroup,
			ValidityStartDate:       data.ValidityStartDate,
			ValidityEndDate:         data.ValidityEndDate,
			CreationDate:            data.CreationDate,
			LastChangeDate:          data.LastChangeDate,
			PlainLongText:           data.PlainLongText,
			IsMarkedForDeletion:     data.IsMarkedForDeletion,
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
			&pm.OpeartionsItem,
			&pm.Sequence,
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.OperationsText,
			&pm.SequenceText,
			&pm.OperationsStatus,
			&pm.ResponsiblePlannerGroup,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.StandardLotSizeQuantity,
			&pm.MinimumLotSizeQuantity,
			&pm.MaximumLotSizeQuantity,
			&pm.OperationsUnit,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.PlainLongText,
			&pm.WorkCenterTypeCode,
			&pm.WorkCenterInternalID,
			&pm.CapacityCategoryCode,
			&pm.OperationCostingRelevancyType,
			&pm.NumberOfTimeTickets,
			&pm.NumberOfConfirmationSlips,
			&pm.OperationSetupType,
			&pm.OperationSetupGroupCategory,
			&pm.OperationSetupGroup,
			&pm.OperationReferenceQuantity,
			&pm.OpQtyToBaseQtyNmrtr,
			&pm.OpQtyToBaseQtyDnmntr,
			&pm.MaximumWaitDuration,
			&pm.MaximumWaitDurationUnit,
			&pm.MinimumWaitDuration,
			&pm.MinimumWaitDurationUnit,
			&pm.StandardQueueDuration,
			&pm.StandardQueueDurationUnit,
			&pm.MinimumQueueDuration,
			&pm.MinimumQueueDurationUnit,
			&pm.StandardMoveDuration,
			&pm.StandardMoveDurationUnit,
			&pm.MinimumMoveDuration,
			&pm.MinimumMoveDurationUnit,
			&pm.OpIsExtlyProcdWithSubcontrg,
			&pm.PlannedDeliveryDuration,
			&pm.Seller,
			&pm.NumberOfOperationPriceUnits,
			&pm.CostElement,
			&pm.OpExternalProcessingPrice,
			&pm.OpExternalProcessingCurrency,
			&pm.OperationScrapPercent,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}

		data := pm
		item = append(item, Item{
			Operations:                    data.Operations,
			OpeartionsItem:                data.OpeartionsItem,
			Sequence:                      data.Sequence,
			Product:                       data.Product,
			BusinessPartner:               data.BusinessPartner,
			Plant:                         data.Plant,
			OperationsText:                data.OperationsText,
			SequenceText:                  data.SequenceText,
			OperationsStatus:              data.OperationsStatus,
			ResponsiblePlannerGroup:       data.ResponsiblePlannerGroup,
			ValidityStartDate:             data.ValidityStartDate,
			ValidityEndDate:               data.ValidityEndDate,
			StandardLotSizeQuantity:       data.StandardLotSizeQuantity,
			MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
			MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
			OperationsUnit:                data.OperationsUnit,
			CreationDate:                  data.CreationDate,
			LastChangeDate:                data.LastChangeDate,
			PlainLongText:                 data.PlainLongText,
			WorkCenterTypeCode:            data.WorkCenterTypeCode,
			WorkCenterInternalID:          data.WorkCenterInternalID,
			CapacityCategoryCode:          data.CapacityCategoryCode,
			OperationCostingRelevancyType: data.OperationCostingRelevancyType,
			NumberOfTimeTickets:           data.NumberOfTimeTickets,
			NumberOfConfirmationSlips:     data.NumberOfConfirmationSlips,
			OperationSetupType:            data.OperationSetupType,
			OperationSetupGroupCategory:   data.OperationSetupGroupCategory,
			OperationSetupGroup:           data.OperationSetupGroup,
			OperationReferenceQuantity:    data.OperationReferenceQuantity,
			OpQtyToBaseQtyNmrtr:           data.OpQtyToBaseQtyNmrtr,
			OpQtyToBaseQtyDnmntr:          data.OpQtyToBaseQtyDnmntr,
			MaximumWaitDuration:           data.MaximumWaitDuration,
			MaximumWaitDurationUnit:       data.MaximumWaitDurationUnit,
			MinimumWaitDuration:           data.MinimumWaitDuration,
			MinimumWaitDurationUnit:       data.MinimumWaitDurationUnit,
			StandardQueueDuration:         data.StandardQueueDuration,
			StandardQueueDurationUnit:     data.StandardQueueDurationUnit,
			MinimumQueueDuration:          data.MinimumQueueDuration,
			MinimumQueueDurationUnit:      data.MinimumQueueDurationUnit,
			StandardMoveDuration:          data.StandardMoveDuration,
			StandardMoveDurationUnit:      data.StandardMoveDurationUnit,
			MinimumMoveDuration:           data.MinimumMoveDuration,
			MinimumMoveDurationUnit:       data.MinimumMoveDurationUnit,
			OpIsExtlyProcdWithSubcontrg:   data.OpIsExtlyProcdWithSubcontrg,
			PlannedDeliveryDuration:       data.PlannedDeliveryDuration,
			Seller:                        data.Seller,
			NumberOfOperationPriceUnits:   data.NumberOfOperationPriceUnits,
			CostElement:                   data.CostElement,
			OpExternalProcessingPrice:     data.OpExternalProcessingPrice,
			OpExternalProcessingCurrency:  data.OpExternalProcessingCurrency,
			OperationScrapPercent:         data.OperationScrapPercent,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
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
