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

type SimplifiedCountryInfo struct {
    Name       string              `json:"name"`
    Continents []string            `json:"continents"`
    Population int                 `json:"population"`
    Languages  map[string]string   `json:"languages"`
    Borders    []string            `json:"borders"`
    Flag       string              `json:"flag"`
    Capital    []string            `json:"capital"`
    Cities     []string            `json:"cities"`
}


type CountryName struct {
    Name struct {
        Common string `json:"common"`
        Official string `json:"official"`
    } `json:"name"`
}

//PopultaionData represents the population for a specific year
type PopulationData struct {
    Year  int `json:"year"`
    Value int `json:"value"`
}

type PopulationResponse struct {
    Mean   int             `json:"mean"`
    Values []PopulationData `json:"values"`
}

type APIData struct {
    Country          string           `json:"country"`
    Code             string           `json:"code"`
    PopulationCounts []PopulationData `json:"populationCounts"`
}

type APIResponse struct {
    Error bool    `json:"error"`
    Msg   string  `json:"msg"`
    Data  APIData `json:"data"`
}

type DiagnosticsResponse struct {
    CountriesNowAPI int    `json:"countriesnowapi"`
    RestCountriesAPI int   `json:"restcountriesapi"`
    Version         string `json:"version"`
    Uptime          int64  `json:"uptime"`
}