package entity

// Constutor
type Model struct {
	Name      string
	Maxtokens int
}

// Classe

func NewModel(name string, maxTokens int) *Model {
	return &Model{
		Name:      name,
		Maxtokens: maxTokens,
	}
}

// Métodos

func (m *Model) GetMaxTokens() int {
	return m.Maxtokens
}

func (m *Model) GetModelName() string {
	return m.Name
}
