package gamehdl

import "github.com/CRISHFAS/Buscaminas-ejemplo-arquitectura-hexagonal/internal/core/domain"

type BodyRevealCell struct {
	Row uint   `json:"row"`
	Col uint   `json:"col"`
}

type ResponseRevealCell domain.Game

func BuildResponseRevealCell(model domain.Game) ResponseRevealCell {
	return ResponseRevealCell(model)
}
