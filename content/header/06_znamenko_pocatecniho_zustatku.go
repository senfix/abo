package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Znaménko počátečního zůstatku – znak + nebo -
type ZnamenkoPocatecnihoZustatku struct {
}

func NewZnamenkoPocatecnihoZustatku() generator.Header {
	return &ZnamenkoPocatecnihoZustatku{}
}

func (t *ZnamenkoPocatecnihoZustatku) Generate(export model.Export) string {
	if export.Header.PocatecniZustatek > 0 {
		return "+"
	}
	return "-"
}

func (t *ZnamenkoPocatecnihoZustatku) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 1, "0", generator.SpacerLeft
}
