# syntax=docker/dockerfile:1

# Building binary
FROM golang:1.19-alpine AS build-stage

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /InfraFolio

# Testing binary
FROM build-stage AS run-test-stage

COPY --from=build-stage /app /app
RUN go test -v ./...

# Deploy to slim image
FROM python:3.11.3-alpine AS build-release-stage

WORKDIR /

COPY scripts/requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY scripts/*.py ./

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

COPY --from=build-stage /InfraFolio /InfraFolio
ENTRYPOINT [ "/InfraFolio" ]

