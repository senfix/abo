package transaction

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Specifický symbol – doplněn vodicími nulami na 10 pozic
//V případě zahraničních plateb bankovní reference
type SpecifickySymbol struct {
}

func NewSpecifickySymbol() generator.Line {
	return &SpecifickySymbol{}
}

func (t *SpecifickySymbol) Generate(header model.Header, transfer model.Transaction) string {
	return transfer.SpecifickySymbol
}

func (t *SpecifickySymbol) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 10, "0", generator.SpacerLeft
}
