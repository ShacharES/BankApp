package util

// Currency constants for supported currencies.
const (
	USD = "USD"
	EUR = "EUR"
	ILS = "ILS"
)

// IsSupportedCurrency checks if the given currency is supported.
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, ILS:
		return true
	default:
		return false
	}
}
