package abo

import "time"

type Transaction struct {
	CisloProtiuctu   string
	CisloDokladu     string
	Castka           float64
	VariabilniSymbol string
	KonstantniSymbol string
	SpecifickySymbol string
	Valuta           time.Time
	DoplnujiciUdaj   string
	Mena             string
	DatumZauctovani  time.Time
}
