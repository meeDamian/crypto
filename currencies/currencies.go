package currencies

import (
	"strings"

	. "github.com/meeDamian/crypto/currencies/symbols"
	"github.com/pkg/errors"
	"github.com/meeDamian/crypto/utils"
)

type Currency struct {
	Name   string
	Symbol string
	IsFiat bool
}

// NOTE: alternative names (aliases) here ONLY. Everything else to `symbols/`
const (
	Bcc = "BCC" // bcash
	Bcg = "BCG" // bgold
	Das = "DAS" // Dash
	Dog = "DOG" // dogecoin
	Drk = "DRK" // Dash…
	Nem = "NEM" // NEM
	Rpx = "RPX" // ripple
	Str = "STR" // Stellar
	Xbt = "XBT" // bitcoin
	Xdg = "XDG" // dogecoin
)

var (
	bcash   = Currency{Bch, "💩", false}
	bgold   = Currency{Btg, "", false}
	bitcoin = Currency{Btc, "₿", false}
	dash    = Currency{Dash, "", false}
	doge    = Currency{Doge, "Ð", false}
	nem     = Currency{Xem, "", false}
	ripple  = Currency{Xrp, "Ʀ", false}
	stellar = Currency{Xlm, "", false}

	list = map[string]Currency{
		// fiat
		Aud: {Aud, "AU$", true},
		Cad: {Cad, "C$", true},
		Clp: {Clp, "C₱", true},
		Cny: {Cny, "C¥", true},
		Cop: {Cop, "CO$", true},
		Eur: {Eur, "€", true},
		Gbp: {Gbp, "£", true},
		Hkd: {Hkd, "HK$", true},
		Idr: {Idr, "Rp", true},
		Jpy: {Jpy, "J¥", true},
		Krw: {Krw, "₩", true},
		Myr: {Myr, "RM", true},
		Ngn: {Ngn, "₦", true},
		Nzd: {Nzd, "NZ$", true},
		Pen: {Pen, "S/", true},
		Php: {Php, "‎₱", true},
		Pln: {Pln, "zł", true},
		Sgd: {Sgd, "SG$", true},
		Thb: {Thb, "‎฿", true},
		Usd: {Usd, "$", true},
		Zar: {Zar, "R", true},

		// crypto (w/more than 1 name)
		Btc:  bitcoin,
		Xbt:  bitcoin,
		Bch:  bcash,
		Bcc:  bcash,
		Doge: doge,
		Dog:  doge,
		Xdg:  doge,
		Bcg:  bgold,
		Btg:  bgold,
		Xrp:  ripple,
		Rpx:  ripple,
		Dash: dash,
		Das:  dash,
		Drk:  dash,
		Nem:  nem,
		Xem:  nem,
		Xlm:  stellar,
		Str:  stellar,

		// crypto
		Aur:   {Aur, "", false},
		Blk:   {Blk, "", false},
		Bts:   {Bts, "", false},
		Cloak: {Cloak, "", false},
		Cure:  {Cure, "", false},
		Efl:   {Efl, "", false},
		Emc2:  {Emc2, "", false},
		Enrg:  {Enrg, "", false},
		Eos:   {Eos, "", false},
		Erc:   {Erc, "", false},
		Etc:   {Etc, "", false},
		Eth:   {Eth, "Ξ", false},
		Evx:   {Evx, "", false},
		Ftc:   {Ftc, "", false},
		Fuel:  {Fuel, "", false},
		Game:  {Game, "", false},
		Gld:   {Gld, "", false},
		Gno:   {Gno, "", false},
		Grs:   {Grs, "", false},
		Hsr:   {Hsr, "", false},
		Hyp:   {Hyp, "", false},
		Icn:   {Icn, "", false},
		Iota:  {Iota, "", false},
		Knc:   {Knc, "", false},
		Kore:  {Kore, "", false},
		Lsk:   {Lsk, "", false},
		Ltc:   {Ltc, "Ł", false},
		Mln:   {Mln, "", false},
		Mona:  {Mona, "", false},
		Neo:   {Neo, "", false},
		Nlg:   {Nlg, "", false},
		Nmc:   {Nmc, "ℕ", false},
		Nxt:   {Nxt, "", false},
		Omg:   {Omg, "", false},
		Pnd:   {Pnd, "", false},
		Pot:   {Pot, "", false},
		Ppc:   {Ppc, "", false},
		Ptc:   {Ptc, "", false},
		Qrk:   {Qrk, "", false},
		Rby:   {Rby, "", false},
		Rdd:   {Rdd, "", false},
		Rep:   {Rep, "", false},
		Slr:   {Slr, "", false},
		Start: {Start, "", false},
		Thc:   {Thc, "", false},
		Trust: {Trust, "", false},
		Ubtc:  {Ubtc, "", false},
		Vrc:   {Vrc, "", false},
		Vtc:   {Vtc, "", false},
		Waves: {Waves, "", false},
		Xcn:   {Xcn, "", false},
		Xdn:   {Xdn, "", false},
		Xmr:   {Xmr, "ɱ", false},
		Xmy:   {Xmy, "", false},
		Xpm:   {Xpm, "", false},
		Xpy:   {Xpy, "", false},
		Xwc:   {Xwc, "", false},
		Xzc:   {Xzc, "", false},
		Zec:   {Zec, "ⓩ", false},

		// other
		Usdt: {Usdt, "US$₮", false},
		Kfee: {Kfee, "ĸ", false},
	}
)

// returns a list of all supported currencies
func All() map[string]Currency {
	return list
}

// returns `Currency` for a supported symbol/alias or error otherwise
func Get(name string) (currency Currency, err error) {
	currency, ok := list[strings.ToUpper(name)]
	if !ok {
		err = errors.Errorf("%s is not a valid currency", name)
	}

	return
}

// returns base name for an alias or unchanged (but uppercase'd) if currency unknown
func Normalise(name string) string {
	currency, err := Get(name)
	if err != nil {
		utils.Log().Debugf("unknown currency %s left unchanged", name)
		return strings.ToUpper(name)
	}

	return currency.Name
}

// returns an alias for a given base name (if found in the `aliases` slice) or unchanged otherwise
func Morph(from string, aliases []string) string {
	for _, alias := range aliases {
		currency, err := Get(alias)
		if err != nil {
			continue
		}

		if currency.Name == from {
			return alias
		}
	}

	return from
}
