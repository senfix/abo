package model

type CurrencyCode string

const (
	AUD CurrencyCode = "00036"
	CAD              = "00124"
	CNY              = "00156"
	CZK              = "00203"
	DKK              = "00208"
	EUR              = "00978"
	GBP              = "00826"
	HRK              = "00191"
	HUF              = "00348"
	CHF              = "00756"
	JPY              = "00392"
	NOK              = "00578"
	PLN              = "00985"
	RON              = "00946"
	RUB              = "00643"
	SEK              = "00752"
	TRY              = "00949"
	USD              = "00840"
)

func NewCurrencyCode(currency string) CurrencyCode {
	switch currency {
	case "AUD":
		return AUD
	case "CAD":
		return CAD
	case "CNY":
		return CNY
	case "CZK":
		return CZK
	case "DKK":
		return DKK
	case "EUR":
		return EUR
	case "GBP":
		return GBP
	case "HRK":
		return HRK
	case "HUF":
		return HUF
	case "CHF":
		return CHF
	case "JPY":
		return JPY
	case "NOK":
		return NOK
	case "PLN":
		return PLN
	case "RON":
		return RON
	case "RUB":
		return RUB
	case "SEK":
		return SEK
	case "TRY":
		return TRY
	case "USD":
		return USD
	}

	panic("unknown currency")
}
