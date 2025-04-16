package responses

import m "dnd-api/db/models"

type RaceResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RacePaginatedResponse struct {
	Data []RaceResponse
	Meta ResponseMeta
}

func NewRaceResponse(race *m.Race) *RaceResponse {
	return &RaceResponse{
		ID:   race.ID,
		Name: race.Name,
	}
}

func NewRaceResponses(races []*m.Race) []RaceResponse {
	var res []RaceResponse
	for _, race := range races {
		res = append(res, *NewRaceResponse(race))
	}
	return res
}

func NewRacePaginatedResponse(races []*m.Race, count, page, pageSize int) *RacePaginatedResponse {
	return &RacePaginatedResponse{
		Data: NewRaceResponses(races),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
