package controller

import (
	"cepclima/internal/entity"
	"cepclima/internal/usecase"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type CepController struct {
	useCase usecase.CepUsecase
}

func NewCepController(useCase usecase.CepUsecase) *CepController {
	return &CepController{useCase: useCase}
}

var cepRegex = regexp.MustCompile(`^[0-9]{5}-?[0-9]{3}$`)

func (cepController *CepController) GetClima(ctx *gin.Context) {
	var input entity.InputCep

	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if !cepRegex.MatchString(input.Cep) {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid zipcode.")
		return
	}

	cepTratado := strings.ReplaceAll(input.Cep, "-", "")

	clima, err := cepController.useCase.GetCepClima(cepTratado)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, clima)
}
