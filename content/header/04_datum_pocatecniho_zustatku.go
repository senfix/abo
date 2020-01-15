package header

import (
	"github.com/senfix/abo/constants"
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Datum počátečního zůstatku ve formátu DDMMRR
type DatumPocatecnihoZustatku struct {
}

func NewDatumPocatecnihoZustatku() generator.Header {
	return &DatumPocatecnihoZustatku{}
}

func (t *DatumPocatecnihoZustatku) Generate(export model.Export) string {
	return export.Header.ExportOd.Format(constants.DEFAULT_TIME_FORMAT)
}

func (t *DatumPocatecnihoZustatku) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 6, "0", generator.SpacerRight
}
