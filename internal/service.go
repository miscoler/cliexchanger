package service

import (
	"context"
	"math"
)

type Converter interface {
	Convert(ctx context.Context, amount float64, srcSymbol, dstSymbol string) (result float64, err error)
}

type exchangeAPI interface {
	GetExchangeRate(ctx context.Context, srcSymbol, dstSymbol string) (rate float64, err error)
}

type converter struct {
	exchanger exchangeAPI
	precision int
}

func (s *converter) Convert(ctx context.Context, amount float64, srcSymbol, dstSymbol string) (result float64, err error) {
	rate, err := s.exchanger.GetExchangeRate(ctx, srcSymbol, dstSymbol)
	if err != nil {
		return
	}
	//basic without correction rule
	result = math.Round(amount*rate*math.Pow10(s.precision)) / math.Pow10(s.precision)
	return
}

func NewConverter(exchanger exchangeAPI, precision int) Converter {
	return &converter{
		exchanger: exchanger,
		precision: precision,
	}
}
