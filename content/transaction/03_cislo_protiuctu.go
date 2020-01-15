package transaction

import (
	"strings"

	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Číslo účtu protistrany – doplněno vodicími nulami na 16 pozic.
//V případě zahraničních (cizoměnových) transakcí je číslo účtu protistrany uvedeno
//ve větě „zpráva pro příjemce“, v tomto poli jsou uvedeny nuly
type CisloProtiUctu struct {
}

func NewCisloProtiUctu() generator.Line {
	return &CisloProtiUctu{}
}

func (t *CisloProtiUctu) Generate(header model.Header, transfer model.Transaction) string {
	cislo := transfer.CisloProtiuctu
	cislo = strings.Replace(cislo, "-", "00", -1)
	index := strings.LastIndex(cislo, "/")
	if index == -1 {
		return cislo
	}
	return cislo[:index]
}
func (t *CisloProtiUctu) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 16, "0", generator.SpacerLeft
}
