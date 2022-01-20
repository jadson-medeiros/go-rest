package models

type Personality struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
}

var Personalities []Personality
