package responses

import (
	m "dnd-api/db/models"
)

type ClassResponse struct {
	ID                  uint    `json:"id"`
	Name                string  `json:"name"`
	ShortDescription    string  `json:"short_description"`
	PrimaryAbility      string  `json:"primary_ability"`
	HitPointDieValue    int     `json:"hit_point_die_value"`
	Saves               string  `json:"saves"`
	SpellcastingAbility *string `json:"spellcasting_ability"`

	Logo            *FileResponse             `json:"logo"`
	BackgroundImage *FileResponse             `json:"background_image"`
	SpellLevels     []ClassSpellLevelResponse `json:"spell_levels"`
}

type ClassPaginatedResponse struct {
	Data []ClassResponse `json:"data" `
	Meta ResponseMeta    `json:"meta"`
}

func NewClassResponse(class *m.Class) *ClassResponse {
	res := &ClassResponse{
		ID:                  class.ID,
		Name:                class.Name,
		ShortDescription:    class.ShortDescription,
		PrimaryAbility:      class.PrimaryAbility,
		HitPointDieValue:    class.HitPointDieValue,
		Saves:               class.Saves,
		SpellcastingAbility: class.SpellcastingAbility,
	}

	if class.Logo.ID != 0 {
		res.Logo = NewFileResponse(&class.Logo)
	}
	if class.BackgroundImage.ID != 0 {
		res.BackgroundImage = NewFileResponse(&class.BackgroundImage)
	}
	if len(class.SpellLevels) != 0 {
		res.SpellLevels = NewCLassSpellLevelResponses(class.SpellLevels)
	}

	return res
}

func NewClassResponses(classes []*m.Class) []ClassResponse {
	var res []ClassResponse
	for _, class := range classes {
		res = append(res, *NewClassResponse(class))
	}
	return res
}

func NewClassPaginatedResponse(classes []*m.Class, count, page, pageSize int) *ClassPaginatedResponse {
	return &ClassPaginatedResponse{
		Data: NewClassResponses(classes),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
