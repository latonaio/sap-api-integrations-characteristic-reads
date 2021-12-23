package responses

type Characteristic struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			DeleteMc                      bool        `json:"Delete_mc"`
			UpdateMc                      bool        `json:"Update_mc"`
			ToCharacteristicDescOc        bool        `json:"to_CharacteristicDesc_oc"`
			ToCharacteristicReferenceOc   bool        `json:"to_CharacteristicReference_oc"`
			ToCharacteristicRestrictionOc bool        `json:"to_CharacteristicRestriction_oc"`
			ToCharacteristicValueOc       bool        `json:"to_CharacteristicValue_oc"`
			CharcInternalID               string      `json:"CharcInternalID"`
			Characteristic                string      `json:"Characteristic"`
			CharcStatus                   string      `json:"CharcStatus"`
			CharcStatusName               string      `json:"CharcStatusName"`
			CharcDataType                 string      `json:"CharcDataType"`
			CharcLength                   int         `json:"CharcLength"`
			CharcDecimals                 int         `json:"CharcDecimals"`
			CharcTemplate                 string      `json:"CharcTemplate"`
			ValueIsCaseSensitive          bool        `json:"ValueIsCaseSensitive"`
			CharcGroup                    string      `json:"CharcGroup"`
			CharcGroupName                string      `json:"CharcGroupName"`
			EntryIsRequired               bool        `json:"EntryIsRequired"`
			MultipleValuesAreAllowed      bool        `json:"MultipleValuesAreAllowed"`
			CharcValueUnit                string      `json:"CharcValueUnit"`
			Currency                      string      `json:"Currency"`
			CharcExponentValue            int         `json:"CharcExponentValue"`
			ValueIntervalIsAllowed        bool        `json:"ValueIntervalIsAllowed"`
			AdditionalValueIsAllowed      bool        `json:"AdditionalValueIsAllowed"`
			NegativeValueIsAllowed        bool        `json:"NegativeValueIsAllowed"`
			ValidityStartDate             string      `json:"ValidityStartDate"`
			ValidityEndDate               string      `json:"ValidityEndDate"`
			ChangeNumber                  string      `json:"ChangeNumber"`
			DocumentType                  string      `json:"DocumentType"`
			DocNumber                     string      `json:"DocNumber"`
			DocumentVersion               string      `json:"DocumentVersion"`
			DocumentPart                  string      `json:"DocumentPart"`
			CharcMaintAuthGrp             string      `json:"CharcMaintAuthGrp"`
			CharcIsReadOnly               bool        `json:"CharcIsReadOnly"`
			CharcIsHidden                 bool        `json:"CharcIsHidden"`
			CharcIsRestrictable           bool        `json:"CharcIsRestrictable"`
			CharcExponentFormat           string      `json:"CharcExponentFormat"`
			CharcEntryIsNotFormatCtrld    bool        `json:"CharcEntryIsNotFormatCtrld"`
			CharcTemplateIsDisplayed      bool        `json:"CharcTemplateIsDisplayed"`
			CreationDate                  string      `json:"CreationDate"`
			LastChangeDate                string      `json:"LastChangeDate"`
			CharcLastChangedDateTime      string      `json:"CharcLastChangedDateTime"`
			KeyDate                       string      `json:"KeyDate"`
			ToCharcDescription struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_CharacteristicDesc"`
			ToValue struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_CharacteristicValue"`
		} `json:"results"`
	} `json:"d"`
}
