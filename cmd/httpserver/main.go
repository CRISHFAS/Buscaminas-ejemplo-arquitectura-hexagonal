package main

import (
	"github.com/CRISHFAS/Buscaminas-ejemplo-arquitectura-hexagonal/internal/core/service/gamesrv"
	"github.com/CRISHFAS/Buscaminas-ejemplo-arquitectura-hexagonal/internal/handlers/gamehdl"
	"github.com/CRISHFAS/Buscaminas-ejemplo-arquitectura-hexagonal/internal/repositories/gamesrepo"
	"github.com/CRISHFAS/Buscaminas-ejemplo-arquitectura-hexagonal/pkg/uidgen"
	"github.com/gin-gonic/gin"
)

func main() {
	gamesRepository := gamesrepo.NewMemKVS()
	gamesService := gamesrv.New(gamesRepository, uidgen.New())
	gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)
	router.PUT("/games/:id", gamesHandler.RevealCell)

	router.Run(":8080")
}
