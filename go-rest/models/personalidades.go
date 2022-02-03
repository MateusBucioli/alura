package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

// Mock de personalidades
var Personalidades []Personalidade
