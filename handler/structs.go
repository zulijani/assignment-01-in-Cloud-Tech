package handler

type CountryInfo struct {
    Name       struct {
        Common   string `json:"common"`
        Official string `json:"official"`
    } `json:"name"`
    Continents []string          `json:"continents"`
    Population int               `json:"population"`
    Languages  map[string]string `json:"languages"`
    Borders    []string          `json:"borders"`
    Flags      struct {
        PNG string `json:"png"`
        SVG string `json:"svg"`
    } `json:"flags"`
    Capital []string `json:"capital"`
    Cities  []string `json:"cities"`
}