package transaction

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Konstantní symbol – doplněn vodicími nulami na 4 pozice
//V případě zahraničních (cizoměnových) transakcí vyplněno nulami
type KonstantniSymbol struct {
}

func NewKonstantniSymbol() generator.Line {
	return &KonstantniSymbol{}
}

func (t *KonstantniSymbol) Generate(header model.Header, transfer model.Transaction) string {
	return transfer.KonstantniSymbol
}

func (t *KonstantniSymbol) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 4, "0", generator.SpacerLeft
}
