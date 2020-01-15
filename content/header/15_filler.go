package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Pořadové číslo výpisu v číslování od začátku roku
type Filler struct {
}

func NewFiller() generator.Header {
	return &Filler{}
}

func (t *Filler) Generate(export model.Export) string {
	return ""
}

func (t *Filler) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 14, " ", generator.SpacerLeft
}
