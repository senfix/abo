package header

import (
	"strings"

	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Číslo účtu včetně předčíslí – doplněno vodicími nulami na 16 pozic
type CisloUctu struct {
}

func NewCisloUctu() generator.Header {
	return &CisloUctu{}
}

func (t *CisloUctu) Generate(export model.Export) string {
	cislo := export.Header.CisloUctu
	cislo = strings.Replace(cislo, "-", "", -1)
	index := strings.LastIndex(cislo, "/")
	return cislo[:index]
}

func (t *CisloUctu) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 16, "0", generator.SpacerLeft
}
