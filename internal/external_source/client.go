package external_source

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/exchangecli/internal/models"

	"github.com/valyala/fasthttp"
)

type Client interface {
	GetExchangeRate(ctx context.Context, srcSymbol, dstSymbol string) (rate float64, err error)
}

type client struct {
	cli         *fasthttp.Client
	exchangeUrl string
}

func (s *client) GetExchangeRate(ctx context.Context, srcSymbol, dstSymbol string) (rate float64, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	req.Header.SetMethod(http.MethodGet)
	req.MultipartForm()
	var reqUri strings.Builder

	reqUri.WriteString(s.exchangeUrl)
	reqUri.WriteString("?base=")
	reqUri.WriteString(srcSymbol)
	reqUri.WriteString("&symbols=")
	reqUri.WriteString(dstSymbol)

	req.SetRequestURI(reqUri.String())
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}

	if res.StatusCode() != http.StatusOK {
		err = errors.New(fmt.Sprintf("%d %s", res.StatusCode(), string(res.Body())))
		return
	}

	var resp models.ExchangeResponse
	err = resp.UnmarshalJSON(res.Body())

	if err != nil {
		return
	}

	var ok bool
	rate, ok = resp.Rates[dstSymbol]

	if !ok {
		err = fmt.Errorf("no exchange rate for %s", dstSymbol)
		return
	}

	return
}

// NewExchangeAPIClient creates new client
func NewExchangeAPIClient(
	exchangeUrl string,
	connectionTimeout time.Duration,
) Client {
	return &client{
		cli: &fasthttp.Client{
			MaxConnDuration: connectionTimeout,
		},
		exchangeUrl: exchangeUrl,
	}
}
