package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Suma kreditních (příchozích) položek – 14 číslic, z toho 12 jsou pozice vyhrazené
//pro celé číslo (včetně vodicích nul) a zbývající 2 pozice jsou desetinná místa (bez
//oddělovače)
type SumaKreditnichPolozek struct {
}

func NewSumaKreditnichPolozek() generator.Header {
	return &SumaKreditnichPolozek{}
}

func (t *SumaKreditnichPolozek) Generate(export model.Export) string {
	sum := 0.0
	for _, t := range export.Transactions {
		if t.Castka < 0 {
			sum += t.Castka
		}
	}
	return generator.FormatMoney(sum)
}

func (t *SumaKreditnichPolozek) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 14, "0", generator.SpacerLeft
}
