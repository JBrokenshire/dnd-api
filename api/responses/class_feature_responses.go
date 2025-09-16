package responses

import m "dnd-api/db/models"

type ClassFeatureResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`

	Options []ClassFeatureOptionResponse `json:"options"`
}

func NewClassFeatureResponse(classFeature *m.ClassFeature) *ClassFeatureResponse {
	res := &ClassFeatureResponse{
		Name:        classFeature.Name,
		Description: classFeature.Description,
		Level:       classFeature.Level,
	}
	if len(classFeature.Options) > 0 {
		res.Options = NewClassFeatureOptionResponses(classFeature.Options)
	}
	return res
}

func NewClassFeatureResponses(classFeatures []*m.ClassFeature) []ClassFeatureResponse {
	var res []ClassFeatureResponse
	for _, classFeature := range classFeatures {
		res = append(res, *NewClassFeatureResponse(classFeature))
	}
	return res
}
