# CEP Clima — Cloud Run

API que recebe um CEP brasileiro e retorna a temperatura atual em Celsius, Fahrenheit e Kelvin.

## Deploy

[https://cloud-run-goexpert-463445218085.us-central1.run.app](https://cloud-run-goexpert-463445218085.us-central1.run.app)

## Rodando localmente com Docker

### Build da imagem

```bash
docker build -t cep-clima .
```

### Subir a aplicação

```bash
docker run cep-clima
```

A API estará disponível em `http://localhost:8080`.

### Exemplo de uso

**Linux/Mac (bash):**
```bash
curl -X GET http://localhost:8080/cep \
  -H "Content-Type: application/json" \
  -d '{"cep":"01310100"}'
```

**Windows (PowerShell):**
```powershell
curl.exe -X GET http://localhost:8080/cep -H "Content-Type: application/json" -d '{\"cep\":\"01310100\"}'
```

## Rodando os testes com Docker

```bash
docker build --target builder -t cep-clima-test .
docker run --rm --entrypoint go cep-clima-test test ./...
```
