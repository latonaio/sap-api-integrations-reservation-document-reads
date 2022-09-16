package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-reservation-document-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			Reservation:                  data.Reservation,
			OrderID:                      data.OrderID,
			GoodsMovementType:            data.GoodsMovementType,
			CostCenter:                   data.CostCenter,
			GoodsRecipientName:           data.GoodsRecipientName,
			ReservationDate:              data.ReservationDate,
			Customer:                     data.Customer,
			WBSElement:                   data.WBSElement,
			ControllingArea:              data.ControllingArea,
			SalesOrder:                   data.SalesOrder,
			SalesOrderItem:               data.SalesOrderItem,
			SalesOrderScheduleLine:       data.SalesOrderScheduleLine,
			AssetNumber:                  data.AssetNumber,
			AssetSubNumber:               data.AssetSubNumber,
			NetworkNumberForAcctAssgmt:   data.NetworkNumberForAcctAssgmt,
			IssuingOrReceivingPlant:      data.IssuingOrReceivingPlant,
			IssuingOrReceivingStorageLoc: data.IssuingOrReceivingStorageLoc,
			ToItem:                       data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			Reservation:                    data.Reservation,
			ReservationItem:                data.ReservationItem,
			RecordType:                     data.RecordType,
			Product:                        data.Product,
			RequirementType:                data.RequirementType,
			MatlCompRequirementDate:        data.MatlCompRequirementDate,
			Plant:                          data.Plant,
			ManufacturingOrderOperation:    data.ManufacturingOrderOperation,
			GoodsMovementIsAllowed:         data.GoodsMovementIsAllowed,
			StorageLocation:                data.StorageLocation,
			Batch:                          data.Batch,
			DebitCreditCode:                data.DebitCreditCode,
			BaseUnit:                       data.BaseUnit,
			GLAccount:                      data.GLAccount,
			GoodsMovementType:              data.GoodsMovementType,
			EntryUnit:                      data.EntryUnit,
			QuantityIsFixed:                data.QuantityIsFixed,
			CompanyCodeCurrency:            data.CompanyCodeCurrency,
			IssuingOrReceivingPlant:        data.IssuingOrReceivingPlant,
			IssuingOrReceivingStorageLoc:   data.IssuingOrReceivingStorageLoc,
			PurchasingDocument:             data.PurchasingDocument,
			PurchasingDocumentItem:         data.PurchasingDocumentItem,
			Supplier:                       data.Supplier,
			ResvnItmRequiredQtyInBaseUnit:  data.ResvnItmRequiredQtyInBaseUnit,
			ReservationItemIsFinallyIssued: data.ReservationItemIsFinallyIssued,
			ReservationItmIsMarkedForDeltn: data.ReservationItmIsMarkedForDeltn,
			ResvnItmRequiredQtyInEntryUnit: data.ResvnItmRequiredQtyInEntryUnit,
			ResvnItmWithdrawnQtyInBaseUnit: data.ResvnItmWithdrawnQtyInBaseUnit,
			ResvnItmWithdrawnAmtInCCCrcy:   data.ResvnItmWithdrawnAmtInCCCrcy,
			GoodsRecipientName:             data.GoodsRecipientName,
			UnloadingPointName:             data.UnloadingPointName,
			ReservationItemText:            data.ReservationItemText,
		})
	}

	return item, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			Reservation:                    data.Reservation,
			ReservationItem:                data.ReservationItem,
			RecordType:                     data.RecordType,
			Product:                        data.Product,
			RequirementType:                data.RequirementType,
			MatlCompRequirementDate:        data.MatlCompRequirementDate,
			Plant:                          data.Plant,
			ManufacturingOrderOperation:    data.ManufacturingOrderOperation,
			GoodsMovementIsAllowed:         data.GoodsMovementIsAllowed,
			StorageLocation:                data.StorageLocation,
			Batch:                          data.Batch,
			DebitCreditCode:                data.DebitCreditCode,
			BaseUnit:                       data.BaseUnit,
			GLAccount:                      data.GLAccount,
			GoodsMovementType:              data.GoodsMovementType,
			EntryUnit:                      data.EntryUnit,
			QuantityIsFixed:                data.QuantityIsFixed,
			CompanyCodeCurrency:            data.CompanyCodeCurrency,
			IssuingOrReceivingPlant:        data.IssuingOrReceivingPlant,
			IssuingOrReceivingStorageLoc:   data.IssuingOrReceivingStorageLoc,
			PurchasingDocument:             data.PurchasingDocument,
			PurchasingDocumentItem:         data.PurchasingDocumentItem,
			Supplier:                       data.Supplier,
			ResvnItmRequiredQtyInBaseUnit:  data.ResvnItmRequiredQtyInBaseUnit,
			ReservationItemIsFinallyIssued: data.ReservationItemIsFinallyIssued,
			ReservationItmIsMarkedForDeltn: data.ReservationItmIsMarkedForDeltn,
			ResvnItmRequiredQtyInEntryUnit: data.ResvnItmRequiredQtyInEntryUnit,
			ResvnItmWithdrawnQtyInBaseUnit: data.ResvnItmWithdrawnQtyInBaseUnit,
			ResvnItmWithdrawnAmtInCCCrcy:   data.ResvnItmWithdrawnAmtInCCCrcy,
			GoodsRecipientName:             data.GoodsRecipientName,
			UnloadingPointName:             data.UnloadingPointName,
			ReservationItemText:            data.ReservationItemText,
		})
	}

	return toItem, nil
}
