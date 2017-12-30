package crypto

type (
	Balance struct {
		Available,
		Total float64
	}

	Balances map[string]Balance
)
