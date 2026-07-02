FROM golang:1.20 AS builder

WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/ai-release-specialist-risk-demo .

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /out/ai-release-specialist-risk-demo /app/ai-release-specialist-risk-demo
EXPOSE 8080
ENTRYPOINT ["/app/ai-release-specialist-risk-demo"]
