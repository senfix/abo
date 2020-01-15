package transaction

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Název protistrany NEBO slovní popis položky
//Poznámka: Pokud je známa protistrana, bude uveden její název pro konkrétní
//transakci, jinak je uveden slovní popis položky.
type DoplnujiciUdaj struct {
}

func NewDoplnujiciUdaj() generator.Line {
	return &DoplnujiciUdaj{}
}

func (t *DoplnujiciUdaj) Generate(header model.Header, transfer model.Transaction) string {
	return transfer.DoplnujiciUdaj
}

func (t *DoplnujiciUdaj) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 20, " ", generator.SpacerRight
}
