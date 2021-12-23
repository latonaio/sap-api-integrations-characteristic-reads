package sap_api_output_formatter

type CharacteristicReads struct {
	ConnectionKey  string `json:"connection_key"`
	Result         bool   `json:"result"`
	RedisKey       string `json:"redis_key"`
	Filepath       string `json:"filepath"`
	APISchema      string `json:"api_schema"`
	Characteristic string `json:"characteristic"`
	Deleted        bool   `json:"deleted"`
}    
    
type Characteristic struct {
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
    ToCharcDescription            string      `json:"to_CharacteristicDesc"`
    ToReference                   string      `json:"to_CharacteristicReference"`
	ToValue                       string      `json:"to_CharacteristicValue"`
}

type CharcDescription struct {
	DeleteMc                 bool        `json:"Delete_mc"`
	UpdateMc                 bool        `json:"Update_mc"`
	CharcInternalID          string      `json:"CharcInternalID"`
	Language                 string      `json:"Language"`
	CharcDescription         string      `json:"CharcDescription"`
	ChangeNumber             string      `json:"ChangeNumber"`
	ValidityStartDate        string      `json:"ValidityStartDate"`
	ValidityEndDate          string      `json:"ValidityEndDate"`
	KeyDate                  string      `json:"KeyDate"`
	CharcLastChangedDateTime string      `json:"CharcLastChangedDateTime"`
}

type ToCharcDescription struct {
	DeleteMc                 bool        `json:"Delete_mc"`
	UpdateMc                 bool        `json:"Update_mc"`
	CharcInternalID          string      `json:"CharcInternalID"`
	Language                 string      `json:"Language"`
	CharcDescription         string      `json:"CharcDescription"`
	ChangeNumber             string      `json:"ChangeNumber"`
	ValidityStartDate        string      `json:"ValidityStartDate"`
	ValidityEndDate          string      `json:"ValidityEndDate"`
	KeyDate                  string      `json:"KeyDate"`
	CharcLastChangedDateTime string      `json:"CharcLastChangedDateTime"`
}

type ToValue struct {
	DeleteMc                  bool        `json:"Delete_mc"`
	UpdateMc                  bool        `json:"Update_mc"`
	ToCharcValueDescOc        bool        `json:"to_CharcValueDesc_oc"`
	CharcInternalID           string      `json:"CharcInternalID"`
	CharcValuePositionNumber  string      `json:"CharcValuePositionNumber"`
	CharcValueDependency      string      `json:"CharcValueDependency"`
	CharcValue                string      `json:"CharcValue"`
	CharcFromNumericValue     string      `json:"CharcFromNumericValue"`
	CharcToNumericValue       string      `json:"CharcToNumericValue"`
	IsDefaultValue            bool        `json:"IsDefaultValue"`
	CharcFromNumericValueUnit string      `json:"CharcFromNumericValueUnit"`
	CharcToNumericValueUnit   string      `json:"CharcToNumericValueUnit"`
	LongTextID                string      `json:"LongTextID"`
	ChangeNumber              string      `json:"ChangeNumber"`
	DocumentType              string      `json:"DocumentType"`
	DocNumber                 string      `json:"DocNumber"`
	DocumentPart              string      `json:"DocumentPart"`
	DocumentVersion           string      `json:"DocumentVersion"`
	ValidityStartDate         string      `json:"ValidityStartDate"`
	ValidityEndDate           string      `json:"ValidityEndDate"`
	KeyDate                   string      `json:"KeyDate"`
	CharcLastChangedDateTime  string      `json:"CharcLastChangedDateTime"`
	ToValueDescription        string      `json:"to_CharcValueDesc"`
}

type ToValueDescription struct {
	DeleteMc                 bool        `json:"Delete_mc"`
	UpdateMc                 bool        `json:"Update_mc"`
	CharcInternalID          string      `json:"CharcInternalID"`
	CharcValuePositionNumber string      `json:"CharcValuePositionNumber"`
	Language                 string      `json:"Language"`
	CharcValueDescription    string      `json:"CharcValueDescription"`
	ChangeNumber             string      `json:"ChangeNumber"`
	ValidityStartDate        string      `json:"ValidityStartDate"`
	ValidityEndDate          string      `json:"ValidityEndDate"`
	KeyDate                  string      `json:"KeyDate"`
	CharcLastChangedDateTime string      `json:"CharcLastChangedDateTime"`
}
