package types

type RequestCity struct {
	Coordinate string `json:"coordinate"`
}

type ResponseCity struct {
	Region   string `json:"region"`
	Province string `json:"province"`
	Regency  string `json:"regency"`
	Name     string `json:"city"`
}
