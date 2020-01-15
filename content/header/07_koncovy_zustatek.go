package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Koncový zůstatek – 14 číslic, z toho 12 jsou pozice vyhrazené pro celé číslo (včetně
//vodicích nul) a zbývající 2 pozice jsou desetinná místa (bez oddělovače). V případě
//měny CZK jde tedy o hodnotu v haléřích
type KoncovyZustatek struct {
}

func NewKoncovyZustatek() generator.Header {
	return &KoncovyZustatek{}
}

func (t *KoncovyZustatek) Generate(export model.Export) string {
	total := export.Header.PocatecniZustatek
	for _, t := range export.Transactions {
		total += t.Castka
	}
	return generator.FormatMoney(total)
}

func (t *KoncovyZustatek) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 14, "0", generator.SpacerLeft
}
