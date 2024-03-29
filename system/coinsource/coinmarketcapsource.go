package coinsource

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/dantas/gocoindesktop/domain"
)

func CoinMarketCapSource() ([]domain.Coin, error) {
	var err error

	marshaledJson, err := fetchMarshaledJson()

	if err != nil {
		return nil, errors.Join(domain.ErrCoinSource, err)
	}

	apiJson, err := parseJson(marshaledJson)

	if err != nil {
		return nil, errors.Join(domain.ErrCoinSource, err)
	}

	var coins = make([]domain.Coin, 0, len(apiJson.Data.CryptoCurrencyList))

	for i := range apiJson.Data.CryptoCurrencyList {
		var coin = &apiJson.Data.CryptoCurrencyList[i]

		if len(coin.Quotes) == 0 {
			continue
		}

		coins = append(coins, domain.Coin{
			Name:  coin.Name,
			Price: coin.Quotes[0].Price,
		})
	}

	return coins, nil
}

const apiUrl = "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing?start=1&limit=50&sortBy=market_cap&sortType=desc&convert=USD&cryptoType=all&tagType=all&audited=false"

func fetchMarshaledJson() ([]byte, error) {
	var response *http.Response
	var err error

	if response, err = http.Get(apiUrl); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	reader := bufio.NewReader(response.Body)

	bytes, err := ioutil.ReadAll(reader)

	return bytes, err
}

func parseJson(marshaledJson []byte) (*jsonFormat, error) {
	var apiJson jsonFormat

	if err := json.Unmarshal(marshaledJson, &apiJson); err != nil {
		return nil, err
	}

	return &apiJson, nil
}

type jsonFormat struct {
	Data struct {
		CryptoCurrencyList []struct {
			Name   string `json:"name"`
			Quotes []struct {
				Price float64 `json:"price"`
			} `json:"quotes"`
		} `json:"cryptoCurrencyList"`
	} `json:"data"`
}
