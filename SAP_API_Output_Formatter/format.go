package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-characteristic-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToCharacteristic(raw []byte, l *logger.Logger) ([]Characteristic, error) {
	pm := &responses.Characteristic{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Characteristic. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	characteristic := make([]Characteristic, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		characteristic = append(characteristic, Characteristic{
	DeleteMc:                      data.DeleteMc,
	UpdateMc:                      data.UpdateMc,
	ToCharacteristicDescOc:        data.ToCharacteristicDescOc,
	ToCharacteristicReferenceOc:   data.ToCharacteristicReferenceOc,
	ToCharacteristicRestrictionOc: data.ToCharacteristicRestrictionOc,
	ToCharacteristicValueOc:       data.ToCharacteristicValueOc,
	CharcInternalID:               data.CharcInternalID,
	Characteristic:                data.Characteristic,
	CharcStatus:                   data.CharcStatus,
	CharcStatusName:               data.CharcStatusName,
	CharcDataType:                 data.CharcDataType,
	CharcLength:                   data.CharcLength,
	CharcDecimals:                 data.CharcDecimals,
	CharcTemplate:                 data.CharcTemplate,
	ValueIsCaseSensitive:          data.ValueIsCaseSensitive,
	CharcGroup:                    data.CharcGroup,
	CharcGroupName:                data.CharcGroupName,
	EntryIsRequired:               data.EntryIsRequired,
	MultipleValuesAreAllowed:      data.MultipleValuesAreAllowed,
	CharcValueUnit:                data.CharcValueUnit,
	Currency:                      data.Currency,
	CharcExponentValue:            data.CharcExponentValue,
	ValueIntervalIsAllowed:        data.ValueIntervalIsAllowed,
	AdditionalValueIsAllowed:      data.AdditionalValueIsAllowed,
	NegativeValueIsAllowed:        data.NegativeValueIsAllowed,
	ValidityStartDate:             data.ValidityStartDate,
	ValidityEndDate:               data.ValidityEndDate,
	ChangeNumber:                  data.ChangeNumber,
	DocumentType:                  data.DocumentType,
	DocNumber:                     data.DocNumber,
	DocumentVersion:               data.DocumentVersion,
	DocumentPart:                  data.DocumentPart,
	CharcMaintAuthGrp:             data.CharcMaintAuthGrp,
	CharcIsReadOnly:               data.CharcIsReadOnly,
	CharcIsHidden:                 data.CharcIsHidden,
	CharcIsRestrictable:           data.CharcIsRestrictable,
	CharcExponentFormat:           data.CharcExponentFormat,
	CharcEntryIsNotFormatCtrld:    data.CharcEntryIsNotFormatCtrld,
	CharcTemplateIsDisplayed:      data.CharcTemplateIsDisplayed,
	CreationDate:                  data.CreationDate,
	LastChangeDate:                data.LastChangeDate,
	CharcLastChangedDateTime:      data.CharcLastChangedDateTime,
	KeyDate:                       data.KeyDate,
	ToCharcDescription:            data.ToCharcDescription.Deferred.URI,
	ToValue:                       data.ToValue.Deferred.URI,
		})
	}

	return characteristic, nil
}

func ConvertToCharcDescription(raw []byte, l *logger.Logger) ([]CharcDescription, error) {
	pm := &responses.CharcDescription{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CharcDescription. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	charcDescription := make([]CharcDescription, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		charcDescription = append(charcDescription, CharcDescription{
	DeleteMc:                 data.DeleteMc,
	UpdateMc:                 data.UpdateMc,
	CharcInternalID:          data.CharcInternalID,
	Language:                 data.Language,
	CharcDescription:         data.CharcDescription,
	ChangeNumber:             data.ChangeNumber,
	ValidityStartDate:        data.ValidityStartDate,
	ValidityEndDate:          data.ValidityEndDate,
	KeyDate:                  data.KeyDate,
	CharcLastChangedDateTime: data.CharcLastChangedDateTime,
		})
	}

	return charcDescription, nil
}

func ConvertToToCharcDescription(raw []byte, l *logger.Logger) ([]ToCharcDescription, error) {
	pm := &responses.ToCharcDescription{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCharcDescription. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toCharcDescription := make([]ToCharcDescription, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toCharcDescription = append(toCharcDescription, ToCharcDescription{
	DeleteMc:                 data.DeleteMc,
	UpdateMc:                 data.UpdateMc,
	CharcInternalID:          data.CharcInternalID,
	Language:                 data.Language,
	CharcDescription:         data.CharcDescription,
	ChangeNumber:             data.ChangeNumber,
	ValidityStartDate:        data.ValidityStartDate,
	ValidityEndDate:          data.ValidityEndDate,
	KeyDate:                  data.KeyDate,
	CharcLastChangedDateTime: data.CharcLastChangedDateTime,
		})
	}

	return toCharcDescription, nil
}

func ConvertToToValue(raw []byte, l *logger.Logger) ([]ToValue, error) {
	pm := &responses.ToValue{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToValue. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toValue := make([]ToValue, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toValue = append(toValue, ToValue{
	DeleteMc:                  data.DeleteMc,
	UpdateMc:                  data.UpdateMc,
	ToCharcValueDescOc:        data.ToCharcValueDescOc,
	CharcInternalID:           data.CharcInternalID,
	CharcValuePositionNumber:  data.CharcValuePositionNumber,
	CharcValueDependency:      data.CharcValueDependency,
	CharcValue:                data.CharcValue,
	CharcFromNumericValue:     data.CharcFromNumericValue,
	CharcToNumericValue:       data.CharcToNumericValue,
	IsDefaultValue:            data.IsDefaultValue,
	CharcFromNumericValueUnit: data.CharcFromNumericValueUnit,
	CharcToNumericValueUnit:   data.CharcToNumericValueUnit,
	LongTextID:                data.LongTextID,
	ChangeNumber:              data.ChangeNumber,
	DocumentType:              data.DocumentType,
	DocNumber:                 data.DocNumber,
	DocumentPart:              data.DocumentPart,
	DocumentVersion:           data.DocumentVersion,
	ValidityStartDate:         data.ValidityStartDate,
	ValidityEndDate:           data.ValidityEndDate,
	KeyDate:                   data.KeyDate,
	CharcLastChangedDateTime:  data.CharcLastChangedDateTime,
	ToValueDescription:        data.ToValueDescription.Deferred.URI,
		})
	}

	return toValue, nil
}

func ConvertToToValueDescription(raw []byte, l *logger.Logger) ([]ToValueDescription, error) {
	pm := &responses.ToValueDescription{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToValueDescription. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toValueDescription := make([]ToValueDescription, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toValueDescription = append(toValueDescription, ToValueDescription{
	DeleteMc:                 data.DeleteMc,
	UpdateMc:                 data.UpdateMc,
	CharcInternalID:          data.CharcInternalID,
	CharcValuePositionNumber: data.CharcValuePositionNumber,
	Language:                 data.Language,
	CharcValueDescription:    data.CharcValueDescription,
	ChangeNumber:             data.ChangeNumber,
	ValidityStartDate:        data.ValidityStartDate,
	ValidityEndDate:          data.ValidityEndDate,
	KeyDate:                  data.KeyDate,
	CharcLastChangedDateTime: data.CharcLastChangedDateTime,
		})
	}

	return toValueDescription, nil
}
