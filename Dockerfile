FROM golang:1.23

WORKDIR /app

COPY .air.toml ./
COPY go.mod go.sum ./
RUN go mod download



RUN go install github.com/air-verse/air@latest && go install github.com/a-h/templ/cmd/templ@latest

COPY . .

EXPOSE 8080 

# Run the air command in the foreground, ensuring the container stays alive
CMD ["air", "-c", ".air.toml"]
