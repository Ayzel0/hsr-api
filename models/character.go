package models

type CharacterJSON struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Types struct {
		PathType struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"pathType"`
		CombatType struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"combatType"`
	} `json:"types"`
}

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Rarity  int    `json:"rank"`
	Path    string
	Element string
}

func (c CharacterJSON) ToCharacter() Character {
	return Character{
		ID:      c.ID,
		Name:    c.Name,
		Rarity:  c.Rank,
		Path:    c.Types.PathType.Name,
		Element: c.Types.CombatType.Name,
	}
}
