package header

import (
	"github.com/senfix/abo/constants"
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Pořadové číslo výpisu v číslování od začátku roku
type DatumVypisu struct {
}

func NewDatumVypisu() generator.Header {
	return &DatumVypisu{}
}

func (t *DatumVypisu) Generate(export model.Export) string {
	return export.Header.ExportDo.Format(constants.DEFAULT_TIME_FORMAT)
}

func (t *DatumVypisu) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 6, "0", generator.SpacerLeft
}
