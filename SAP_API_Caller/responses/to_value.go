package responses

type ToValue struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
			ToValueDescription        struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_CharcValueDesc"`
		} `json:"results"`
	} `json:"d"`
}
