package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockExchangeApi struct {
	mock.Mock
}

func (m *mockExchangeApi) GetExchangeRate(ctx context.Context,
	srcSymbol, dstSymbol string) (rate float64, err error) {
	args := m.Called(srcSymbol, dstSymbol)
	return args.Get(0).(float64), args.Error(1)
}

func TestConvert(t *testing.T) {
	exchange := &mockExchangeApi{}

	exchange.On("GetExchangeRate", "RUB", "USD").Return(79.23, nil)
	exchange.On("GetExchangeRate", "RUB", "RUB").Return(1.0, nil)

	type fields struct {
		exchanger exchangeAPI
		precision int
	}
	type args struct {
		ctx       context.Context
		amount    float64
		srcSymbol string
		dstSymbol string
		rate      func(methodName string, arguments ...interface{}) mock.Call
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult float64
		wantErr    bool
	}{
		{
			name: "RUB USD",
			fields: fields{
				exchanger: exchange,
				precision: 4,
			},
			args: args{
				ctx:       context.Background(),
				amount:    2345.4,
				srcSymbol: "RUB",
				dstSymbol: "USD",
			},
			wantResult: 185826.042,
			wantErr:    false,
		}, {
			name: "RUB RUB",
			fields: fields{
				exchanger: exchange,
				precision: 4,
			},
			args: args{
				ctx:       context.Background(),
				amount:    2345.4,
				srcSymbol: "RUB",
				dstSymbol: "RUB",
			},
			wantResult: 2345.4,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewConverter(
				tt.fields.exchanger,
				tt.fields.precision,
			)
			gotResult, err := s.Convert(tt.args.ctx, tt.args.amount, tt.args.srcSymbol, tt.args.dstSymbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("Convert() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
