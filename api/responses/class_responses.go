package responses

import m "dnd-api/db/models"

type ClassResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ClassPaginatedResponse struct {
	Data []ClassResponse
	Meta ResponseMeta
}

func NewClassResponse(class *m.Class) *ClassResponse {
	return &ClassResponse{
		ID:   class.ID,
		Name: class.Name,
	}
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
