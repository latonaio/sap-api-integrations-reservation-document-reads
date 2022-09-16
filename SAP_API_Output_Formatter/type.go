package sap_api_output_formatter

type ReservationDocument struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	Reservation   string `json:"reservation_document"`
	Deleted       bool   `json:"deleted"`
}

type Header struct {
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
	ToItem                       string `json:"to_ReservationDocumentItem"`
}

type Item struct {
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
}

type ToItem struct {
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
}
