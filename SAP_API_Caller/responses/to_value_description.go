package responses

type ToValueDescription struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
