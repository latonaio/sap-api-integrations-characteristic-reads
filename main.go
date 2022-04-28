package main

import (
	sap_api_caller "sap-api-integrations-characteristic-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-characteristic-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Characteristic_Charc_Description_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Characteristic", "CharcDescription",
		}
	}

	caller.AsyncGetCharacteristic(
		inoutSDC.Characteristic.Characteristic,
		inoutSDC.Characteristic.CharcDescription.Language,
		inoutSDC.Characteristic.CharcDescription.CharcDescription,

		accepter,
	)
}
