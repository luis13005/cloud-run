package service

import (
	"cepclima/internal/entity"
	"context"
	"encoding/json"
	"net/http"
	"os"
)

const (
	cepUrl    = `https://viacep.com.br/ws/`
	cepUrlEnd = `/json/`
	climaUrl  = `http://api.weatherapi.com/v1/current.json`
)

type CepClimaService struct{}

func NewCepClimaService() *CepClimaService {
	return &CepClimaService{}
}

func (s *CepClimaService) GetClimaByCep(cidade string) (*entity.ClimaOutput, error) {
	weatherApi := os.Getenv("WEATHER_API")

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, climaUrl, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", weatherApi)
	q.Add("q", cidade)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var clima entity.Clima
	json.NewDecoder(res.Body).Decode(&clima)

	climaOutput := &entity.ClimaOutput{Celsius: clima.Current.Temp_c, Fahrenheit: clima.Current.Temp_f, Kelvin: clima.Current.Temp_c + 273}
	return climaOutput, nil
}

func (s *CepClimaService) GetCep(cepString string) (*entity.Cep, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, cepUrl+cepString+cepUrlEnd, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cep entity.Cep

	json.NewDecoder(resp.Body).Decode(&cep)
	return &cep, nil
}

func (s *CepClimaService) GetClima(cidade string) (*entity.ClimaOutput, error) {

	return nil, nil
}
