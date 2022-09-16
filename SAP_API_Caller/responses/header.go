package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Reservation                  string `json:"Reservation"`
			OrderID                      string `json:"OrderID"`
			GoodsMovementType            string `json:"GoodsMovementType"`
			CostCenter                   string `json:"CostCenter"`
			GoodsRecipientName           string `json:"GoodsRecipientName"`
			ReservationDate              string `json:"ReservationDate"`
			Customer                     string `json:"Customer"`
			WBSElement                   string `json:"WBSElement"`
			ControllingArea              string `json:"ControllingArea"`
			SalesOrder                   string `json:"SalesOrder"`
			SalesOrderItem               string `json:"SalesOrderItem"`
			SalesOrderScheduleLine       string `json:"SalesOrderScheduleLine"`
			AssetNumber                  string `json:"AssetNumber"`
			AssetSubNumber               string `json:"AssetSubNumber"`
			NetworkNumberForAcctAssgmt   string `json:"NetworkNumberForAcctAssgmt"`
			IssuingOrReceivingPlant      string `json:"IssuingOrReceivingPlant"`
			IssuingOrReceivingStorageLoc string `json:"IssuingOrReceivingStorageLoc"`
			ToItem                       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ReservationDocumentItem"`
		} `json:"results"`
	} `json:"d"`
}
