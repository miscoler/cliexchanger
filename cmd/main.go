package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	service "github.com/exchangecli/internal"
	"github.com/valyala/fasthttp"

	"github.com/exchangecli/internal/external_source"
)

const rateURLConst = "https://api.exchangeratesapi.io/latest"
const precisionConst = 4

func main() {

	if len(os.Args) < 4 {
		fmt.Println("not enough arguments")
		os.Exit(0)
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	srcSymbol := os.Args[2]
	dstSymbol := os.Args[3]

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	exchanger := external_source.NewExchangeAPIClient(rateURLConst, fasthttp.DefaultDialTimeout)

	converter := service.NewConverter(exchanger, precisionConst)

	result, err := converter.Convert(context.Background(), amount, srcSymbol, dstSymbol)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Printf("%.2f %s equals %.2f %s", amount, srcSymbol, result, dstSymbol)
}
