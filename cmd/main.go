package main

import (
	"cepclima/internal/infra/controller"
	"cepclima/internal/infra/service"
	"cepclima/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()
	cepController := dependencias()

	r.POST("/cep", cepController.GetClima)

	r.Run(":8080")
}

func dependencias() *controller.CepController {
	if err := godotenv.Load("cmd/.env"); err != nil {
		log.Fatalf("Erro ao carregar váriaveis de ambiente")
	}

	svc := service.NewCepClimaService()
	cepUseCase := usecase.NewCepUsecase(svc)
	cepController := controller.NewCepController(*cepUseCase)
	return cepController
}
