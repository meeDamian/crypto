package acx

import (
	"encoding/json"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

// NOTE: this is just a playground material for now

const myTradesUrl = "https://acx.io/api/v2/trades/my.json"

type trade struct {
	Id       int64     `json:"id"`
	Price    string    `json:"price"`
	Volume   string    `json:"volume"`
	Funds    string    `json:"funds"`
	Market   string    `json:"market"`
	PlacedAt time.Time `json:"created_at"`
	Trend    string    `json:"trend"`
	Side     string    `json:"side"`
	OrderId  int64     `json:"order_id"`
}

func MyTrades(c crypto.Credentials) (ts []trade, err error) {
	res, err := privateRequest(c, "GET", myTradesUrl, map[string]string{
		"market": "btcaud",
		"limit":  "1000",
	})

	if err != nil {
		return []trade{}, err
	}

	log.Debugln(res.Status)

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ts)
	if err != nil {
		err = errors.Wrapf(err, "can't decode own trades from %s", Domain)
	}

	return
}

//func MyOrders(c crypto.Credentials) (o []crypto.Order, err error) {
//	resp, err := privateRequest(c, "GET", "", map[string]string{
//		"market": "btcaud",
//	})
//	if err != nil {
//		return []crypto.Order{}, err
//	}
//
//	defer resp.Body.Close()
//
//	respBody, _ := ioutil.ReadAll(resp.Body)
//
//	log.Println(resp.Status)
//	log.Println(string(respBody))
//
//	return
//}
