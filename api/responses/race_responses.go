package responses

import m "dnd-api/db/models"

type RaceResponse struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	CreatureType     string `json:"creature_type"`
	Size             string `json:"size"`
	BaseSpeed        int    `json:"base_speed"`

	Logo   *FileResponse   `json:"logo"`
	Traits []TraitResponse `json:"traits"`
}

type RacePaginatedResponse struct {
	Data []RaceResponse `json:"data"`
	Meta ResponseMeta   `json:"meta"`
}

func NewRaceResponse(race *m.Race) *RaceResponse {
	res := &RaceResponse{
		ID:               race.ID,
		Name:             race.Name,
		ShortDescription: race.ShortDescription,
		CreatureType:     race.CreatureType,
		Size:             race.Size,
		BaseSpeed:        race.BaseSpeed,
	}
	if race.Logo.ID != 0 {
		res.Logo = NewFileResponse(&race.Logo)
	}
	if len(race.Traits) > 0 {
		res.Traits = NewTraitResponses(race.Traits)
	}

	return res
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
