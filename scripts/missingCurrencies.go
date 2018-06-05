package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/exchanges"
	"github.com/meeDamian/crypto/utils"
)

var minimumOccurrences = 7

func main() {
	utils.EnableDebugLogs(true)

	missing := make(map[string]int)
	var missingCount int

	crypto.SetCurrencyNotSupportedTrigger(func(currency string, err ...error) {
		if err != nil {
			utils.Log().Error(err)
			return
		}

		c := strings.ToUpper(currency)
		missing[c]++
		missingCount++
	})

	for _, e := range exchanges.All {
		log := utils.Log().WithField("exchange", e.Domain)
		_, err := e.Markets()
		if err != nil {
			log.Error(err)
		}
	}

	log := utils.Log()

	log.Println(missing, "total:", missingCount)

	var symbols, currs string
	for symbol, cnt := range missing {
		if cnt < minimumOccurrences {
			log.Printf("skipping %s, because it exists in less than %d markets only\n", symbol, minimumOccurrences)
			continue
		}

		if _, err := strconv.Atoi(symbol[:1]); err == nil {
			log.Printf("skipping %s, because it starts with a number\n", symbol)
			continue
		}

		bla := fmt.Sprintf("%s%s", strings.ToUpper(symbol[:1]), strings.ToLower(symbol[1:]))
		symbols += fmt.Sprintf("%s = \"%s\" // NAME(%d)         URL\n", bla, symbol, cnt)
		currs += fmt.Sprintf("%s:{%s, \"\", false},\n", bla, bla)
	}

	log.Printf("symbols.go:\n%s\n", symbols)
	log.Printf("currencies.go:\n%s\n", currs)

	log.Println("NOTE: Adding above and running this script again might result in more missing currencies being detected!")
}
