package transaction

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Variabilní symbol – doplněn vodicími nulami na 10 pozic
//V případě zahraničních (cizoměnových) transakcí vyplněno nulami
type VariabilniSymbol struct {
}

func NewVariabilniSymbol() generator.Line {
	return &VariabilniSymbol{}
}

func (t *VariabilniSymbol) Generate(header model.Header, transfer model.Transaction) string {
	return transfer.VariabilniSymbol
}

func (t *VariabilniSymbol) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 10, "0", generator.SpacerLeft
}
