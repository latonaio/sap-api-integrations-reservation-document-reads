package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Reservation                    string `json:"Reservation"`
			ReservationItem                string `json:"ReservationItem"`
			RecordType                     string `json:"RecordType"`
			Product                        string `json:"Product"`
			RequirementType                string `json:"RequirementType"`
			MatlCompRequirementDate        string `json:"MatlCompRequirementDate"`
			Plant                          string `json:"Plant"`
			ManufacturingOrderOperation    string `json:"ManufacturingOrderOperation"`
			GoodsMovementIsAllowed         bool   `json:"GoodsMovementIsAllowed"`
			StorageLocation                string `json:"StorageLocation"`
			Batch                          string `json:"Batch"`
			DebitCreditCode                string `json:"DebitCreditCode"`
			BaseUnit                       string `json:"BaseUnit"`
			GLAccount                      string `json:"GLAccount"`
			GoodsMovementType              string `json:"GoodsMovementType"`
			EntryUnit                      string `json:"EntryUnit"`
			QuantityIsFixed                bool   `json:"QuantityIsFixed"`
			CompanyCodeCurrency            string `json:"CompanyCodeCurrency"`
			IssuingOrReceivingPlant        string `json:"IssuingOrReceivingPlant"`
			IssuingOrReceivingStorageLoc   string `json:"IssuingOrReceivingStorageLoc"`
			PurchasingDocument             string `json:"PurchasingDocument"`
			PurchasingDocumentItem         string `json:"PurchasingDocumentItem"`
			Supplier                       string `json:"Supplier"`
			ResvnItmRequiredQtyInBaseUnit  string `json:"ResvnItmRequiredQtyInBaseUnit"`
			ReservationItemIsFinallyIssued bool   `json:"ReservationItemIsFinallyIssued"`
			ReservationItmIsMarkedForDeltn bool   `json:"ReservationItmIsMarkedForDeltn"`
			ResvnItmRequiredQtyInEntryUnit string `json:"ResvnItmRequiredQtyInEntryUnit"`
			ResvnItmWithdrawnQtyInBaseUnit string `json:"ResvnItmWithdrawnQtyInBaseUnit"`
			ResvnItmWithdrawnAmtInCCCrcy   string `json:"ResvnItmWithdrawnAmtInCCCrcy"`
			GoodsRecipientName             string `json:"GoodsRecipientName"`
			UnloadingPointName             string `json:"UnloadingPointName"`
			ReservationItemText            string `json:"ReservationItemText"`
		} `json:"results"`
	} `json:"d"`
}
