package abo

import (
	"bytes"
	"time"

	"github.com/senfix/abo/model"

	"github.com/senfix/abo/service"
)

type Generator interface {
	AddTransaction(tx Transaction)
	Generate() []byte
	GetEndBalance() float64
}

type generator struct {
	export model.Export
}

func New(
	cisloUctu string,
	nazevUctu string,
	pocatecniZustatek float64,
	exportOd time.Time,
	exportDo time.Time,
) Generator {
	export := model.Export{
		Header: model.Header{
			CisloUctu:         cisloUctu,
			NazevUctu:         nazevUctu,
			PocatecniZustatek: pocatecniZustatek,
			ExportOd:          exportOd,
			ExportDo:          exportDo,
		},
	}
	return &generator{
		export: export,
	}
}
func (t *generator) AddTransaction(tx Transaction) {
	direction := model.Incoming
	if tx.Castka < 0 {
		direction = model.Outgoing
	}

	t.export.Transactions = append(t.export.Transactions, model.Transaction{
		CisloProtiuctu:   tx.CisloProtiuctu,
		CisloDokladu:     tx.CisloDokladu,
		Castka:           tx.Castka,
		Direction:        direction,
		VariabilniSymbol: tx.VariabilniSymbol,
		KonstantniSymbol: tx.KonstantniSymbol,
		SpecifickySymbol: tx.SpecifickySymbol,
		Valuta:           tx.Valuta,
		DoplnujiciUdaj:   tx.DoplnujiciUdaj,
		KodMeny:          model.NewCurrencyCode(tx.Mena),
		DatumZauctovani:  tx.DatumZauctovani,
	})
}

func (t *generator) Generate() []byte {
	var data []byte
	writter := bytes.NewBuffer(data)

	headerPrinter := service.HeaderPrinter{}
	transferPrinter := service.TransactionLinePrinter{}

	headerPrinter.Export(t.export, writter)
	for _, tx := range t.export.Transactions {
		transferPrinter.Export(t.export.Header, tx, writter)
	}

	return writter.Bytes()
}
func (t *generator) GetEndBalance() float64 {
	total := t.export.Header.PocatecniZustatek
	for _, t := range t.export.Transactions {
		total += t.Castka
	}
	return total
}
