# Étape 1 : Construction de l'application
FROM golang:1.25 AS builder

# Définit le répertoire de travail
WORKDIR /app

# Copie les fichiers de dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copie le reste du code
COPY . .

# Construit l'application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /server .

# Étape 2 : Image finale légère
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /server /server

EXPOSE 8080

CMD ["/server"]
