package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-operations-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-operations-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemOperation *[]dpfm_api_output_formatter.ItemOperation
	var itemOperationComponent *[]dpfm_api_output_formatter.ItemOperationComponent
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeaderByOwnerProductionPlantBP":
			func() {
				header = c.HeaderByOwnerProductionPlantBP(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()
		case "ItemOperation":
			func() {
				itemOperation = c.ItemOperation(mtx, input, output, errs, log)
			}()
		case "ItemOperations":
			func() {
				itemOperation = c.ItemOperations(mtx, input, output, errs, log)
			}()
		case "ItemOperationComponent":
			func() {
				itemOperationComponent = c.ItemOperationComponent(mtx, input, output, errs, log)
			}()
		case "ItemOperationComponents":
			func() {
				itemOperationComponent = c.ItemOperationComponents(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                 header,
		Item:                   item,
		ItemOperation:          itemOperation,
		ItemOperationComponent: itemOperationComponent,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	operations := input.Header.Operations

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_header_data
		WHERE Operations = ?;`, operations,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderByOwnerProductionPlantBP(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.OwnerProductionPlantBusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND OwnerProductionPlantBusinessPartner = %v", where, *input.Header.OwnerProductionPlantBusinessPartner)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_header_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	operations := input.Header.Operations
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		args = append(args, operations, v.OperationsItem)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_item_data
		WHERE (Operations, OpeartionsItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Items(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	item := &dpfm_api_input_reader.Item{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	where := fmt.Sprintf("WHERE Operations = %v", input.Header.Operations)
	if item != nil {
		if item.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *item.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_item_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperation {
	var args []interface{}
	operations := input.Header.Operations
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			args = append(args, operations, v.OperationsItem, w.OperationID)
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_item_operation_data
		WHERE (Operations, OpeartionsItem, OperationID) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperation {
	itemOperation := &dpfm_api_input_reader.ItemOperation{}
	if len(input.Header.Item) > 0 && len(input.Header.Item[0].ItemOperation) > 0 {
		itemOperation = &input.Header.Item[0].ItemOperation[0]
	}
	where := fmt.Sprintf("WHERE Operations = %v", input.Header.Operations)
	if input.Header.Item != nil {
		where = fmt.Sprintf("%s\nAND OperationsItem = %v", where, input.Header.Item[0].OperationsItem)
	}
	if itemOperation != nil {
		if itemOperation.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *itemOperation.IsMarkedForDeletion)
		}
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_item_operation_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperationComponent(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperationComponent {
	var args []interface{}
	cnt := 0
	for _, v := range input.Header.Item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			itemOperationComponent := w.ItemOperationComponent
			for _, y := range itemOperationComponent {
				args = append(args, w.Operations, w.OperationsItem, w.OperationID, v.BillOfMaterial, y.BillOfMaterialItem)
			}
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?,?),", cnt-1) + "(?,?,?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data-platform-operations-sql-item-operation-component-data
		WHERE (Operations, OperationsItem, OperationID, BillOfMaterial, BillOfMaterialItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperationComponent(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	return data
}

func (c *DPFMAPICaller) ItemOperationComponents(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperationComponent {
	item := &dpfm_api_input_reader.Item{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	where := fmt.Sprintf("WHERE Operations = %v", input.Header.Operations)
	if item != nil {
		if item.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *item.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_item_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperationComponent(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
