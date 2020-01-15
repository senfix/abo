package transaction

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Kód účtování – 1 pro debetní (odchozí) položku, 2 pro kreditní (příchozí) položku,
//4 pro storno debetní položky a 5 pro storno kreditní položky
type KodUctovani struct {
}

func NewKodUctovani() generator.Line {
	return &KodUctovani{}
}

func (t *KodUctovani) Generate(header model.Header, transfer model.Transaction) string {
	switch transfer.Direction {
	case model.Incoming:
		return "2"
	case model.Outgoing:
		return "1"
	}
	return "1"
}

func (t *KodUctovani) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 1, "0", generator.SpacerRight
}
