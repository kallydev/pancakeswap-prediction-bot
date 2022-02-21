# Pancakeswap prediction bot

![GitHub last commit](https://img.shields.io/github/last-commit/kallydev/pancakeswap-prediction-bot?style=flat-square)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/kallydev/pancakeswap-prediction-bot?style=flat-square)
![GitHub license](https://img.shields.io/github/license/kallydev/pancakeswap-prediction-bot?style=flat-square)

## Introduction

This bot will bet the specified amount of BNB to the side with more money at the last N seconds of the bet countdown.

This is because I chose to follow the direction of the public considering that usually no one loses money on purpose, but of course this can be disturbed by other bots.

## Warning

This robot is not guaranteed to profit you and there are bound to be losers in this game!

This strategy is not stable under extreme conditions. In testing, it managed to double my BNB, then dropped back to 1.2x, and finally increased to 3x until I open-sourced the code with only 2x profit.

I have tested a total of 400 rounds, my win rate is around 70%.

## Build

```shell
# Ubuntu or Debian
sudo apt update
sudo apt install git snapd upx build-essential musl-dev -y
sudo snap install go --classic

git clone https://github.com/kallydev/pancakeswap-prediction-bot
cd pancakeswap-prediction-bot
make build && ls ./build
```

## Usage

### Screen

```shell
# Linux
screen -S pancakeswap-prediction-bot
./build/pancakeswap-prediction-bot_linux_amd64 --help
./build/pancakeswap-prediction-bot_linux_amd64 --private-key replace_me_as_your_private_key --amount 0.01

# Ctrl + A + D
```

### Docker

```shell
make build_image

docker run \
  --name pancakeswap-prediction-bot \
  --restart=on-failure:3 \
  -dit kallydev/pancakeswap-prediction-bot:latest \
  --private-key replace_me_as_your_private_key \
  --amount 0.01
```

## Donate

### ETH

You can donate any amount to me in the Ethereum `Mainnet`, `Polygon` or `BEP20` to support my work.

```diff
+ 0x000000A52a03835517E9d193B3c27626e1Bc96b1
```

## License

[GNU General Public License v3.0](LICENSE).
