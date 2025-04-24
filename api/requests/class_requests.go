package requests

type CreateClassRequest struct {
	Name             string `json:"name" validate:"required,max=200" example:"Barbarian"`
	ShortDescription string `json:"short_description" validate:"required" example:"Barbarians are mighty warriors who are powered by primal forces of the multiverse that manifest as a Rage."`
	PrimaryAbility   string `json:"primary_ability" validate:"required,max=200" example:"Strength"`
	HitPointDieValue int    `json:"hit_point_die_value" validate:"required" example:"4"`
	Saves            string `json:"saves" validate:"required,max=200" example:"['Strength','Constitution']"`
}

type UpdateClassRequest struct {
	Name             string `json:"name" validate:"required,max=200" example:"Barbarian"`
	ShortDescription string `json:"short_description" validate:"required" example:"Barbarians are mighty warriors who are powered by primal forces of the multiverse that manifest as a Rage."`
	PrimaryAbility   string `json:"primary_ability" validate:"required,max=200" example:"Strength"`
	HitPointDieValue int    `json:"hit_point_die_value" validate:"required" example:"4"`
	Saves            string `json:"saves" validate:"required,max=200" example:"['Strength','Constitution']"`
}
