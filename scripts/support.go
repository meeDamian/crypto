package main

import (
	"fmt"
	"strings"

	"github.com/meeDamian/crypto/exchanges"
)

func main() {
	header := fmt.Sprintf("| %-15s | %-10s | %-12s | %-17s | %-10s", "Domain", ".Markets()", ".OrderBook()", ".AllOrderBooks()", ".Balances()")
	length := len(header)

	fmt.Println(header)
	fmt.Println(strings.Repeat("-", length+1))

	var markets, obs, aobs, balances int64
	for _, e := range exchanges.All {
		fmt.Printf("| %-15s | %-10t | %-12t | %-17t | %t\n", e.Domain, e.Markets != nil, e.OrderBook != nil, e.AllOrderBooks != nil, e.Balances != nil)

		if e.Markets != nil {
			markets++
		}

		if e.OrderBook != nil {
			obs++
		}

		if e.AllOrderBooks != nil {
			aobs++
		}

		if e.Balances != nil {
			balances++
		}
	}

	fmt.Println(strings.Repeat("-", length+1))
	fmt.Printf("| %15s | %-10d | %-12d | %-17d | %d", "done:", markets, obs, aobs, balances)
}
