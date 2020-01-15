package transaction

import (
	"strings"

	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Číslo účtu včetně předčíslí – doplněno vodicími nulami na 16 pozic
type CisloUctu struct {
}

func NewCisloUctu() generator.Line {
	return &CisloUctu{}
}

func (t *CisloUctu) Generate(header model.Header, transfer model.Transaction) string {
	cislo := header.CisloUctu
	cislo = strings.Replace(cislo, "-", "", -1)
	index := strings.LastIndex(cislo, "/")
	return cislo[:index]
}

func (t *CisloUctu) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 16, "0", generator.SpacerLeft
}
