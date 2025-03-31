package responses

import m "dnd-api/db/models"

type ClassResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ClassPaginatedResponse struct {
	Data []ClassResponse `json:"data"`
	Meta ResponseMeta    `json:"meta"`
}

func NewClassResponse(class *m.Class) *ClassResponse {
	return &ClassResponse{
		ID:   class.ID,
		Name: class.Name,
	}
}

func NewClassResponses(classes []m.Class) []ClassResponse {
	res := make([]ClassResponse, 0)
	for _, class := range classes {
		res = append(res, *NewClassResponse(&class))
	}

	return res
}

func NewClassPaginatedResponse(classes []m.Class, count, page, pageSize int) *ClassPaginatedResponse {
	return &ClassPaginatedResponse{
		Data: NewClassResponses(classes),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
