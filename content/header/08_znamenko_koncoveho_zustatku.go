package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Znaménko koncového zůstatku – znak + nebo -
type ZnamenkoKoncovehoZustatku struct {
}

func NewZnamenkoKoncovehoZustatku() generator.Header {
	return &ZnamenkoKoncovehoZustatku{}
}

func (t *ZnamenkoKoncovehoZustatku) Generate(export model.Export) string {
	total := export.Header.PocatecniZustatek
	for _, t := range export.Transactions {
		total += t.Castka
	}
	if total > 0 {
		return "+"
	}
	return "-"
}

func (t *ZnamenkoKoncovehoZustatku) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 1, " ", generator.SpacerLeft
}
