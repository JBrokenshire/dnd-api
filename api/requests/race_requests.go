package requests

type CreateRaceRequest struct {
	Name             string `json:"name" validate:"required,max=200" example:"Dragonborn"`
	ShortDescription string `json:"short_description" validate:"required" example:"The ancestors of dragonborn hatched from the eggs of chromatic and metallic dragons."`
	CreatureType     string `json:"creature_type" validate:"required,max=200" example:"Humanoid"`
	Size             string `json:"size" validate:"required,max=200" example:"Medium (about 5-7 feet tall)"`
	BaseSpeed        int    `json:"base_speed" validate:"required" example:"30"`
}

type UpdateRaceRequest struct {
	Name             string `json:"name" validate:"required,max=200" example:"Dragonborn"`
	ShortDescription string `json:"short_description" validate:"required" example:"The ancestors of dragonborn hatched from the eggs of chromatic and metallic dragons."`
	CreatureType     string `json:"creature_type" validate:"required,max=200" example:"Humanoid"`
	Size             string `json:"size" validate:"required,max=200" example:"Medium (about 5-7 feet tall)"`
	BaseSpeed        int    `json:"base_speed" validate:"required" example:"30"`
}
