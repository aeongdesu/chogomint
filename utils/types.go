package utils

type Beatmap struct {
	Checksum string `json:"checksum"`
	Id int `json:"id"`
}

type Beatmaps struct {
	Id int `json:"id"`
}

type Collection struct {
	Name string `json:"name"`
	Beatmapsets []Beatmaps `json:"beatmapsets"`
	Id int `json:"id"`
}