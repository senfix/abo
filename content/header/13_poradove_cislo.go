package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Pořadové číslo výpisu v číslování od začátku roku
type PoradoveCislo struct {
}

func NewPoradoveCislo() generator.Header {
	return &PoradoveCislo{}
}

func (t *PoradoveCislo) Generate(export model.Export) string {
	return export.Header.ExportOd.Format("01")
}

func (t *PoradoveCislo) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 3, "0", generator.SpacerLeft
}
