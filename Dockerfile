FROM golang:1.21.3 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -C ./ ./main.go


FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/main .

EXPOSE 8080

USER nonroot:nonroot

CMD ["./main"]