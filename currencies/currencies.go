package currencies

import (
	"strings"

	. "github.com/meeDamian/crypto/currencies/symbols"
	"github.com/pkg/errors"
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
		Inr: {Inr, "₹", true},
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
		Ada:   {Ada, "", false},
		Adt:   {Adt, "", false},
		Adx:   {Adx, "", false},
		Agrs:  {Agrs, "", false},
		Ant:   {Ant, "", false},
		Aur:   {Aur, "", false},
		Bat:   {Bat, "", false},
		Bcy:   {Bcy, "", false},
		Blk:   {Blk, "", false},
		Block: {Block, "", false},
		Bnt:   {Bnt, "", false},
		Bsd:   {Bsd, "", false},
		Bts:   {Bts, "", false},
		Byc:   {Byc, "", false},
		Cfi:   {Cfi, "", false},
		Cloak: {Cloak, "", false},
		Crb:   {Crb, "", false},
		Cure:  {Cure, "", false},
		Cvc:   {Cvc, "", false},
		Dgb:   {Dgb, "", false},
		Dgd:   {Dgd, "", false},
		Dnt:   {Dnt, "", false},
		Dope:  {Dope, "", false},
		Dtb:   {Dtb, "", false},
		Efl:   {Efl, "", false},
		Emc2:  {Emc2, "", false},
		Emc:   {Emc, "", false},
		Eng:   {Eng, "", false},
		Enrg:  {Enrg, "", false},
		Eos:   {Eos, "", false},
		Erc:   {Erc, "", false},
		Etc:   {Etc, "", false},
		Eth:   {Eth, "Ξ", false},
		Evx:   {Evx, "", false},
		Fct:   {Fct, "", false},
		Flo:   {Flo, "", false},
		Ftc:   {Ftc, "", false},
		Fuel:  {Fuel, "", false},
		Fun:   {Fun, "", false},
		Game:  {Game, "", false},
		Gbg:   {Gbg, "", false},
		Gld:   {Gld, "", false},
		Gno:   {Gno, "", false},
		Gnt:   {Gnt, "", false},
		Grs:   {Grs, "", false},
		Gup:   {Gup, "", false},
		Hmq:   {Hmq, "", false},
		Hsr:   {Hsr, "", false},
		Hyp:   {Hyp, "", false},
		Icn:   {Icn, "", false},
		Infx:  {Infx, "", false},
		Iop:   {Iop, "", false},
		Iota:  {Iota, "", false},
		Knc:   {Knc, "", false},
		Kore:  {Kore, "", false},
		Lgd:   {Lgd, "", false},
		Lsk:   {Lsk, "", false},
		Ltc:   {Ltc, "Ł", false},
		Lun:   {Lun, "", false},
		Maid:  {Maid, "", false},
		Mana:  {Mana, "", false},
		Mco:   {Mco, "", false},
		Meme:  {Meme, "", false},
		Mer:   {Mer, "", false},
		Mln:   {Mln, "", false},
		Mona:  {Mona, "", false},
		Mtl:   {Mtl, "", false},
		Mue:   {Mue, "", false},
		Myst:  {Myst, "", false},
		Nav:   {Nav, "", false},
		Neo:   {Neo, "", false},
		Nlg:   {Nlg, "", false},
		Nmc:   {Nmc, "ℕ", false},
		Nmr:   {Nmr, "", false},
		Nxt:   {Nxt, "", false},
		Omg:   {Omg, "", false},
		Omni:  {Omni, "", false},
		Part:  {Part, "", false},
		Pay:   {Pay, "", false},
		Pivx:  {Pivx, "", false},
		Pnd:   {Pnd, "", false},
		Pot:   {Pot, "", false},
		Powr:  {Powr, "", false},
		Ppc:   {Ppc, "", false},
		Ptc:   {Ptc, "", false},
		Ptoy:  {Ptoy, "", false},
		Qash:  {Qash, "", false},
		Qrk:   {Qrk, "", false},
		Qrl:   {Qrl, "", false},
		Qtum:  {Qtum, "", false},
		Rby:   {Rby, "", false},
		Rcn:   {Rcn, "", false},
		Rdd:   {Rdd, "", false},
		Rep:   {Rep, "", false},
		Rlc:   {Rlc, "", false},
		Salt:  {Salt, "", false},
		Sc:    {Sc, "", false},
		Slr:   {Slr, "", false},
		Sls:   {Sls, "", false},
		Snt:   {Snt, "", false},
		Srn:   {Srn, "", false},
		Start: {Start, "", false},
		Storj: {Storj, "", false},
		Strat: {Strat, "", false},
		Thc:   {Thc, "", false},
		Tix:   {Tix, "", false},
		Trst:  {Trst, "", false},
		Trust: {Trust, "", false},
		Trx:   {Trx, "", false},
		Ubtc:  {Ubtc, "", false},
		Ukg:   {Ukg, "", false},
		Vib:   {Vib, "", false},
		Vrc:   {Vrc, "", false},
		Vrm:   {Vrm, "", false},
		Vtc:   {Vtc, "", false},
		Vtr:   {Vtr, "", false},
		Waves: {Waves, "", false},
		Wings: {Wings, "", false},
		Xcn:   {Xcn, "", false},
		Xdn:   {Xdn, "", false},
		Xel:   {Xel, "", false},
		Xmr:   {Xmr, "ɱ", false},
		Xmy:   {Xmy, "", false},
		Xpm:   {Xpm, "", false},
		Xpy:   {Xpy, "", false},
		Xvg:   {Xvg, "", false},
		Xwc:   {Xwc, "", false},
		Xzc:   {Xzc, "", false},
		Zcl:   {Zcl, "", false},
		Zec:   {Zec, "ⓩ", false},

		// other
		Kfee: {Kfee, "ĸ", false},
		Usdt: {Usdt, "US$₮", false},
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
		err = errors.Errorf("unknown currency '%s'", name)
	}

	return
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
