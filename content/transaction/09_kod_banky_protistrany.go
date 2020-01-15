package transaction

import (
	"strings"

	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Kód banky protistrany (např. 0300)
//V případě zahraničních (cizoměnových) transakcí vyplněno nulami, SWIFT kód banky
//je uveden ve větě „zpráva pro příjemce“
type KodBankyProtistrany struct {
}

func NewKodBankyProtistrany() generator.Line {
	return &KodBankyProtistrany{}
}

func (t *KodBankyProtistrany) Generate(header model.Header, transfer model.Transaction) string {
	cislo := transfer.CisloProtiuctu
	index := strings.LastIndex(cislo, "/")
	return cislo[index+1:]
}

func (t *KodBankyProtistrany) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 4, "0", generator.SpacerRight
}
