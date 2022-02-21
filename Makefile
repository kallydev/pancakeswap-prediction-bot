.PHONY: build
build:
	mkdir -p ./build
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ./build/pancakeswap-prediction-bot_linux_amd64 ./command/pancakeswap.go
	upx -9 ./build/pancakeswap-prediction-bot_linux_amd64
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o ./build/pancakeswap-prediction-bot_windows_amd64.exe ./command/pancakeswap.go
	upx -9 ./build/pancakeswap-prediction-bot_windows_amd64.exe
	GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o ./build/pancakeswap-prediction-bot_darwin_amd64 ./command/pancakeswap.go
	upx -9 ./build/pancakeswap-prediction-bot_darwin_amd64

build_image:
	docker build -t kallydev/pancakeswap-prediction-bot:v0.1.0 -t kallydev/pancakeswap-prediction-bot:latest .
