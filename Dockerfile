FROM golang:1.17-alpine AS build

WORKDIR /root/pancakeswap-prediction-bot

COPY . .

RUN apk add --no-cache gcc musl-dev

RUN mkdir -p ./build && go build -ldflags "-w -s" -o ./build/pancakeswap-prediction-bot ./command/pancakeswap.go

FROM alpine:3.15

WORKDIR /root/pancakeswap-prediction-bot

COPY --from=build /root/pancakeswap-prediction-bot/build/pancakeswap-prediction-bot /root/pancakeswap-prediction-bot/build/pancakeswap-prediction-bot

ENTRYPOINT ["/root/pancakeswap-prediction-bot/build/pancakeswap-prediction-bot"]
