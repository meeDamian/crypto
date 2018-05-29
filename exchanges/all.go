package exchanges

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/exchanges/acx"
	"github.com/meeDamian/crypto/exchanges/binance"
	"github.com/meeDamian/crypto/exchanges/bitbay"
	"github.com/meeDamian/crypto/exchanges/bitcoincoid"
	"github.com/meeDamian/crypto/exchanges/bitfinex"
	"github.com/meeDamian/crypto/exchanges/bitstamp"
	"github.com/meeDamian/crypto/exchanges/bittrex"
	"github.com/meeDamian/crypto/exchanges/bx"
	"github.com/meeDamian/crypto/exchanges/coinfloor"
	"github.com/meeDamian/crypto/exchanges/gdax"
	"github.com/meeDamian/crypto/exchanges/hitbtc"
	"github.com/meeDamian/crypto/exchanges/kraken"
	"github.com/meeDamian/crypto/exchanges/liqui"
	"github.com/meeDamian/crypto/exchanges/luno"
	"github.com/meeDamian/crypto/exchanges/nzbcx"
	"github.com/meeDamian/crypto/exchanges/quoinex"
	"github.com/meeDamian/crypto/exchanges/surbtc"
	"github.com/meeDamian/crypto/exchanges/tdax"
	"github.com/meeDamian/crypto/exchanges/yobit"
)

var All = []crypto.Exchange{
	acx.New(),
	binance.New(),
	bitbay.New(),
	bitcoincoid.New(),
	bitfinex.New(),
	bitstamp.New(),
	bittrex.New(),
	bx.New(),
	coinfloor.New(),
	gdax.New(),
	hitbtc.New(),
	kraken.New(),
	liqui.New(),
	luno.New(),
	nzbcx.New(),
	quoinex.New(),
	surbtc.New(),
	tdax.New(),
	yobit.New(),
}
