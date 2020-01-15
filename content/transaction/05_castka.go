package transaction

import (
	"fmt"
	"math"

	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Částka zaúčtované transakce – 14 číslic, z toho 12 jsou pozice vyhrazené pro celé
//číslo (včetně vodicích nul) a zbývající 2 pozice jsou desetinná místa (bez oddělovače)
//Poznámka: Pro měnu CZK jde o částku v haléřích.
type Castka struct {
}

func NewCastka() generator.Line {
	return &Castka{}
}

func (t *Castka) Generate(header model.Header, transfer model.Transaction) string {
	castka := transfer.Castka
	castka = math.Floor(castka * 100)
	castka = math.Abs(castka)
	return fmt.Sprintf("%.0f", castka)
}

func (t *Castka) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 12, "0", generator.SpacerLeft
}
