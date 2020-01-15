package transaction

import (
	"github.com/senfix/abo/constants"
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Datum valuty ve formátu DDMMRR (standardně shodné s datem zaúčtování)
type Valuta struct {
}

func NewValuta() generator.Line {
	return &Valuta{}
}

func (t *Valuta) Generate(header model.Header, transfer model.Transaction) string {
	return transfer.Valuta.Format(constants.DEFAULT_TIME_FORMAT)
}

func (t *Valuta) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 6, " ", generator.SpacerRight
}
