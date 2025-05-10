package ports

import "github.com/CRISHFAS/Buscaminas-ejemplo-arquitectura-hexagonal/internal/core/domain"

type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}
