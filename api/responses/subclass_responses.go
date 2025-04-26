package responses

import (
	m "dnd-api/db/models"
)

type SubclassResponse struct {
	ID               uint   `json:"id"`
	ClassId          uint   `json:"class_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`

	Logo *FileResponse `json:"logo"`
}

type SubclassPaginatedResponse struct {
	Data []SubclassResponse `json:"data"`
	Meta ResponseMeta       `json:"meta"`
}

func NewSubclassResponse(subclass *m.Subclass) *SubclassResponse {
	res := &SubclassResponse{
		ID:               subclass.ID,
		ClassId:          subclass.ClassId,
		Name:             subclass.Name,
		ShortDescription: subclass.ShortDescription,
	}

	if subclass.Logo.ID != 0 {
		res.Logo = NewFileResponse(&subclass.Logo)
	}

	return res
}

func NewSubclassResponses(subclasses []*m.Subclass) []SubclassResponse {
	var res []SubclassResponse
	for _, subclass := range subclasses {
		res = append(res, *NewSubclassResponse(subclass))
	}
	return res
}

func NewSubclassPaginatedResponse(subclasses []*m.Subclass, count, page, pageSize int) *SubclassPaginatedResponse {
	return &SubclassPaginatedResponse{
		Data: NewSubclassResponses(subclasses),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
