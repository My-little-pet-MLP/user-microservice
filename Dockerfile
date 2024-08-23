# Etapa de construção
FROM golang:1.23.0-alpine as builder

WORKDIR /app
COPY . .

# Construir o binário da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o /user-microservice .
# Etapa final
FROM scratch

WORKDIR /app

# Copiar o binário compilado da etapa de construção
COPY --from=builder /user-microservice /user-microservice

# Copiar o arquivo .env se ele existir
# COPY --from=builder /app/.env /app/.env

EXPOSE 8080

# Comando para executar a aplicação
CMD ["/user-microservice"]