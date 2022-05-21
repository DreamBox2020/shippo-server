package model

type Picture struct {
	Model
	Path    string `json:"path"`
	Uri     string `json:"url"`
	Name    string `json:"name"`
	Mime    string `json:"mime"`
	Channel string `json:"channel"`
}
