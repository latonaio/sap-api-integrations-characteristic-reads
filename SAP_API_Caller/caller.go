package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-characteristic-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetCharacteristic(characteristic, charcLanguage, charcDescription string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Characteristic":
			func() {
				c.Characteristic(characteristic)
				wg.Done()
			}()
		case "CharcDescription":
			func() {
				c.CharcDescription(charcLanguage, charcDescription)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Characteristic(characteristic string) {
	charcData, err := c.callCharacteristicSrvAPIRequirementCharacteristic("A_ClfnCharacteristicForKeyDate", characteristic)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(charcData)

	valueData, err := c.callToValue(charcData[0].ToValue)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(valueData)

	charcDescriptionData, err := c.callToCharcDescription(charcData[0].ToCharcDescription)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(charcDescriptionData)
	
	valueDescriptionData, err := c.callToValueDescription(valueData[0].ToValueDescription)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(valueDescriptionData)

}

func (c *SAPAPICaller) callCharacteristicSrvAPIRequirementCharacteristic(api, characteristic string) ([]sap_api_output_formatter.Characteristic, error) {
	url := strings.Join([]string{c.baseURL, "API_CLFN_CHARACTERISTIC_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCharacteristic(req, characteristic)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCharacteristic(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToValue(url string) ([]sap_api_output_formatter.ToValue, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToValue(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCharcDescription(url string) ([]sap_api_output_formatter.ToCharcDescription, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCharcDescription(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToValueDescription(url string) ([]sap_api_output_formatter.ToValueDescription, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToValueDescription(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CharcDescription(charcLanguage, charcDescription string) {
	data, err := c.callCharacteristicSrvAPIRequirementCharcDescription("A_ClfnCharcDescForKeyDate", charcLanguage, charcDescription)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callCharacteristicSrvAPIRequirementCharcDescription(api, charcLanguage, charcDescription string) ([]sap_api_output_formatter.CharcDescription, error) {
	url := strings.Join([]string{c.baseURL, "API_CLFN_CHARACTERISTIC_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCharcDescription(req, charcLanguage, charcDescription)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCharcDescription(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithCharacteristic(req *http.Request, characteristic string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Characteristic eq '%s'", characteristic))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithCharcDescription(req *http.Request, charcLanguage, charcDescription string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Language eq '%s' and substringof('%s', CharcDescription)", charcLanguage, charcDescription))
	req.URL.RawQuery = params.Encode()
}
