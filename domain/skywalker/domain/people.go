package domain

type People struct {
	PeopleCharacter []*PeopleCharacter `json:"results"`
}

type PeopleCharacter struct {
	Name  string   `json:"name"`
	Films []string `json:"films"`
}
