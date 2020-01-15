package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Počáteční zůstatek – 14 číslic, z toho 12 jsou pozice vyhrazené pro celé číslo (včetně
//vodicích nul) a zbývající 2 pozice jsou desetinná místa (bez oddělovače). V případě
//měny CZK jde tedy o hodnotu v haléřích
type PocatecniZustatek struct {
}

func NewPocatecniZustatek() generator.Header {
	return &PocatecniZustatek{}
}

func (t *PocatecniZustatek) Generate(export model.Export) string {
	return generator.FormatMoney(export.Header.PocatecniZustatek)
}

func (t *PocatecniZustatek) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 14, "0", generator.SpacerLeft
}
