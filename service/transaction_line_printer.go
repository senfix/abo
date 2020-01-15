package service

import (
	"io"

	"github.com/senfix/abo/constants"

	tx "github.com/senfix/abo/content/transaction"
	"github.com/senfix/abo/model"

	"github.com/senfix/abo/generator"
)

//https://www.csob.cz/portal/documents/10710/1927786/format-gpc.pdf
type TransactionLinePrinter struct {
}

func (t *TransactionLinePrinter) Export(header model.Header, transaction model.Transaction, writer io.Writer) {

	generators := []generator.Line{
		tx.NewTypZaznamu(),
		tx.NewCisloUctu(),
		tx.NewCisloProtiUctu(),
		tx.NewCisloDokladu(),
		tx.NewCastka(),
		tx.NewKodUctovani(),
		tx.NewVariabilniSymbol(),
		tx.NewSpacer(),
		tx.NewKodBankyProtistrany(),
		tx.NewKonstantniSymbol(),
		tx.NewSpecifickySymbol(),
		tx.NewValuta(),
		tx.NewDoplnujiciUdaj(),
		tx.NewKodMeny(),
		tx.NewDatumZauctovani(),
	}

	for _, g := range generators {
		part := g.Generate(header, transaction)
		writer.Write([]byte(sanitize(g, part)))
	}
	writer.Write([]byte(constants.LINE_SEPARATOR))

}
