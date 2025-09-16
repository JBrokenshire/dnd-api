package responses

import m "dnd-api/db/models"

type ClassFeatureOptionResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewClassFeatureOptionResponse(classFeatureOption *m.ClassFeatureOption) *ClassFeatureOptionResponse {
	return &ClassFeatureOptionResponse{
		Name:        classFeatureOption.Name,
		Description: classFeatureOption.Description,
	}
}

func NewClassFeatureOptionResponses(classFeatureOptions []*m.ClassFeatureOption) []ClassFeatureOptionResponse {
	var res []ClassFeatureOptionResponse
	for _, classFeatureOption := range classFeatureOptions {
		res = append(res, *NewClassFeatureOptionResponse(classFeatureOption))
	}
	return res
}
