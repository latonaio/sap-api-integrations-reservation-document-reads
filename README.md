# sap-api-integrations-reservation-document-reads
sap-api-integrations-reservation-document-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 入出庫予定データを取得するマイクロサービスです。    
sap-api-integrations-reservation-document-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-reservation-document-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_RESERVATION_DOCUMENT_SRV_0001/overview  

## 動作環境  
sap-api-integrations-reservation-document-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-reservation-document-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-reservation-document-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_RESERVATION_DOCUMENT_SRV_0001/overview    
* APIサービス名(=baseURL): API_RESERVATION_DOCUMENT_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-reservation-document-reads には、次の API をコールするためのリソースが含まれています。  

* A_ReservationDocumentHeader（入出庫予定 - ヘッダ）
* A_ReservationDocumentItem（入出庫予定 - 明細）
* ToItem（入出庫予定 - 明細 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-reservation-document-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.Reservation.Reservation（入出庫予定）
* inoutSDC.Reservation.ReservationItem.RecordType（レコードタイプ）
* inoutSDC.Reservation.ReservationItem.Product（品目）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。

```
	"api_schema": "/A_ReservationDocumentHeader",
	"accepter": ["Header"],
	"reservation_document": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "/A_ReservationDocumentHeader",
	"accepter": ["All"],
	"reservation_document": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetReservationDocument(reservation, recordType, product string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(reservation)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(reservation, recordType, product)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 入出庫予定 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"Reservation" ～ "to_ReservationDocumentItem" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-reservation-document-reads/SAP_API_Caller/caller.go#L60",
	"function": "sap-api-integrations-reservation-document-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"Reservation": "1",
			"OrderID": "",
			"GoodsMovementType": "201",
			"CostCenter": "100000",
			"GoodsRecipientName": "Test Accept",
			"ReservationDate": "2022-09-16",
			"Customer": "",
			"WBSElement": "",
			"ControllingArea": "0001",
			"SalesOrder": "",
			"SalesOrderItem": "0",
			"SalesOrderScheduleLine": "0",
			"AssetNumber": "",
			"AssetSubNumber": "",
			"NetworkNumberForAcctAssgmt": "",
			"IssuingOrReceivingPlant": "",
			"IssuingOrReceivingStorageLoc": "",
			"to_ReservationDocumentItem": "http://100.21.57.120:8080/sap/opu/odata/sap/API_RESERVATION_DOCUMENT_SRV/A_ReservationDocumentHeader('1')/to_ReservationDocumentItem"
		}
	],
	"time": "2022-09-16T23:52:16+09:00"
}

```
