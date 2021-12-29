package model

type ResponsePokemon struct {
	Count   int      `json:"count"`
	Results []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ResponseDetailPokemon struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Types   []Types `json:"types"`
	Moves   []Moves `json:"moves"`
	Sprites Sprites `json:"sprites"`
}

type Types struct {
	Slot int         `json:"slot"`
	Type interface{} `json:"type"`
}

type Moves struct {
	Move interface{} `json:"move"`
}

type Sprites struct {
	BackDefault string `json:"back_default"`
}
