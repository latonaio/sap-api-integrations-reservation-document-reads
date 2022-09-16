package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-reservation-document-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

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

func (c *SAPAPICaller) Header(reservation string) {
	headerData, err := c.callReservationDocumentSrvAPIRequirementHeader("A_ReservationDocumentHeader", reservation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)
}

func (c *SAPAPICaller) callReservationDocumentSrvAPIRequirementHeader(api, reservation string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_RESERVATION_DOCUMENT_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, reservation)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(reservation, recordType, product string) {
	data, err := c.callReservationDocumentSrvAPIRequirementItem("A_ReservationDocumentItem", reservation, recordType, product)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callReservationDocumentSrvAPIRequirementItem(api, reservation, recordType, product string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_RESERVATION_DOCUMENT_SRV", api}, "/")

	param := c.getQueryWithItem(map[string]string{}, reservation, recordType, product)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, reservation string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Reservation eq '%s'", reservation)
	return params
}

func (c *SAPAPICaller) getQueryWithItem(params map[string]string, reservation, recordType, product string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Reservation eq '%s' and RecordType eq '%s' and Product eq '%s'", reservation, recordType, product)
	return params
}
