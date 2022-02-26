package domain

type Response struct {
	Characters []*Character
}

type Character struct {
	Name  string
	Films []string
}

func NewCharacter(name string, films []string) *Character {
	return &Character{
		Name:  name,
		Films: films,
	}
}
