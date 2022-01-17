package models

type Personality struct {
	Name  string `json:"name"`
	About string `json:"about"`
}

var Personalities []Personality
