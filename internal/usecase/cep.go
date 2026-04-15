package usecase

import (
	"cepclima/internal/entity"
	"errors"
)

type CepUsecase struct {
	service ExternalServiceInterface
}

func NewCepUsecase(service ExternalServiceInterface) *CepUsecase {
	return &CepUsecase{
		service: service,
	}
}

type ExternalServiceInterface interface {
	GetClimaByCep(cidade string) (*entity.ClimaOutput, error)

	GetCep(cep string) (*entity.Cep, error)
}

type UseCaseInterface interface {
	GetCepClima(cep string) (*entity.ClimaOutput, error)
}

func (useCase *CepUsecase) GetCepClima(cepString string) (*entity.ClimaOutput, error) {

	cep, err := useCase.service.GetCep(cepString)
	if err != nil {
		return nil, errors.New("can not find zipcode")
	}

	climaOut, err := useCase.service.GetClimaByCep(cep.Localidade)
	if err != nil {
		return nil, err
	}

	return climaOut, nil
}
