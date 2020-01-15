package transaction

import (
	"github.com/senfix/abo/constants"
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Datum zaúčtování ve formátu DDMMRR
type DatumZauctovani struct {
}

func NewDatumZauctovani() generator.Line {
	return &DatumZauctovani{}
}

func (t *DatumZauctovani) Generate(header model.Header, transfer model.Transaction) string {
	return transfer.DatumZauctovani.Format(constants.DEFAULT_TIME_FORMAT)
}

func (t *DatumZauctovani) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 6, " ", generator.SpacerRight
}
