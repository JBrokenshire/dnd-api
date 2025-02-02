package responses

import m "github.com/JBrokenshire/dnd-api/db/models"

type ClassResponse struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
}

type ClassPaginatedResponse struct {
	Data []ClassResponse `json:"data"`
	Meta ResponseMeta    `json:"meta"`
}

func NewClassResponse(class *m.Class) ClassResponse {
	return ClassResponse{
		ID:               class.ID,
		Name:             class.Name,
		ShortDescription: class.ShortDescription,
		LongDescription:  class.LongDescription,
	}
}

func NewClassResponses(classes []*m.Class) []ClassResponse {
	var res []ClassResponse
	for _, class := range classes {
		res = append(res, NewClassResponse(class))
	}
	return res
}

func NewClassPaginatedResponse(classes []*m.Class, totalCount, page, pageSize int) ClassPaginatedResponse {
	return ClassPaginatedResponse{
		Data: NewClassResponses(classes),
		Meta: ResponseMeta{TotalCount: totalCount, Page: page, PageSize: pageSize},
	}
}
