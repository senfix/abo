package model

import "time"

type Transaction struct {
	CisloProtiuctu      string
	CisloDokladu        string
	Castka              float64
	Direction           TransactionDirection
	VariabilniSymbol    string
	KodBankyProtistrany string
	KonstantniSymbol    string
	SpecifickySymbol    string
	Valuta              time.Time
	DoplnujiciUdaj      string
	KodMeny             CurrencyCode
	DatumZauctovani     time.Time
}
