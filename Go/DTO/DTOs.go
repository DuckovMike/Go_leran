package DTOs

type GoblinDTO struct {
	Name   string `json:"name"`
	Health int    `json:"health"`
	Attack int    `json:"attack"`
}

func NewGoblin(n string, h int, a int) GoblinDTO {
	return GoblinDTO{
		Name:   n,
		Health: h,
		Attack: a,
	}
}
