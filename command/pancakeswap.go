package main

import (
	"context"
	"github.com/kallydev/pancakeswap-prediction-bot/internal/bot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func main() {
	config := &bot.Config{}
	command := &cobra.Command{
		Use: "pancakeswap-prediction-bot",
		Run: func(cmd *cobra.Command, args []string) {
			predictionBot, err := bot.New(context.Background(), config)
			if err != nil {
				logrus.Fatalln(err)
			}
			if err := predictionBot.Run(); err != nil {
				logrus.Fatalln(err)
			}
		},
	}
	command.Flags().StringVar(&config.Endpoint, "endpoint", "https://bsc-dataseed.binance.org/", "")
	command.Flags().StringVar(&config.PrivateKeyHex, "private-key", "", "")
	command.Flags().Float64Var(&config.Amount, "amount", 0, "")
	command.Flags().IntVar(&config.Time, "time", 10, "")
	if err := command.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}
