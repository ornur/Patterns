package main

type Hero struct {
	Name string
	Type string
}

func (h *Hero) getName() string {
	return h.Name
}
func (h *Hero) getType() string {
	return h.Type
}

func (h *Hero) setName(name string) {
	h.Name = name
}

func (h *Hero) setType(heroType string) {
	h.Type = heroType
}
