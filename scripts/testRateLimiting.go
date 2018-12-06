package main

import (
	markets2 "github.com/meeDamian/crypto/markets"
	"net/http"
	"sync"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/exchanges"
	"github.com/meeDamian/crypto/utils"
)

const (
	// set to 0 to use all available markets, set to anything else to hardcode the amount of requests
	marketsCount = 0

	// set to ex: acx.Domain to only run on one exchange
	exchangeName = ""
)

// _Go_ through all exchanges and determine what's a possible rate limitation
func main() {
	for _, exchange := range exchanges.All {
		if exchangeName != "" && exchange.Domain != exchangeName {
			continue
		}

		log := utils.Log().WithField("exchange", exchange.Domain)

		start := time.Now()

		markets, err := exchange.Markets()
		if err != nil {
			log.Errorln("couldn't fetch markets: ", err)
			continue
		}

		if marketsCount != 0 {
			for len(markets) <= marketsCount {
				log.Printf("too few(%d) - multiplying", len(markets))
				markets = append(markets, markets...)
			}

			if len(markets) > marketsCount {
				log.Printf("too many(%d) - cutting", len(markets))
				markets = markets[:marketsCount]
			}
		}

		var wg sync.WaitGroup
		wg.Add(len(markets))

		var success, rateLimited, errored int64

		for _, m := range markets {
			go func(exchange crypto.Exchange, m markets2.Market) {
				defer func() {
					if r := recover(); r != nil {
						log.Errorf("skipping %s due to: %v", m, r)
					}

					wg.Done()
				}()

				ob, err := exchange.OrderBook(m)
				if err != nil {
					log.Errorf("couldn't download %s ob: %v", m, err)

					if err.Error() == http.StatusText(http.StatusTooManyRequests) {
						rateLimited++
						return
					}

					errored++
					return
				}

				success++
				log.Infof("Downloaded %s: %d|%d", m, len(ob.Bids), len(ob.Asks))

			}(exchange, m)
		}

		wg.Wait()

		log.WithField("took", time.Now().Sub(start)).Infof("Stats: %d ok, %d rate limited, %d error'd", success, rateLimited, errored)
	}
}
