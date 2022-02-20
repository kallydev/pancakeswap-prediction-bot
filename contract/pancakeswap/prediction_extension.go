package pancakeswap

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

type Round struct {
	Epoch               *big.Int
	StartTimestamp      *big.Int
	LockTimestamp       *big.Int
	CloseTimestamp      *big.Int
	LockPrice           *big.Int
	ClosePrice          *big.Int
	LockOracleId        *big.Int
	CloseOracleId       *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}

func (_Prediction *PredictionCaller) Round(opts *bind.CallOpts, arg0 *big.Int) (*Round, error) {
	round, err := _Prediction.Rounds(opts, arg0)
	if err != nil {
		return nil, err
	}
	return &Round{
		Epoch:               round.Epoch,
		StartTimestamp:      round.StartTimestamp,
		LockTimestamp:       round.LockTimestamp,
		CloseTimestamp:      round.CloseTimestamp,
		LockPrice:           round.LockPrice,
		ClosePrice:          round.ClosePrice,
		LockOracleId:        round.LockOracleId,
		CloseOracleId:       round.CloseOracleId,
		TotalAmount:         round.TotalAmount,
		BullAmount:          round.BullAmount,
		BearAmount:          round.BearAmount,
		RewardBaseCalAmount: round.RewardBaseCalAmount,
		RewardAmount:        round.RewardAmount,
		OracleCalled:        round.OracleCalled,
	}, nil
}
