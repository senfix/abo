package transaction

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Identifikátor transakce na 13 pozic, přičemž:
//•	 pozice 36 (první číslice) určuje typ transakce – 0 pro příkazy vzešlé z příkazu
//k inkasu a 1 pro ostatní platby
//•	 pozice 43–48 (posledních 6 cifer) určuje sekvenční číslo v historii účtu
//V případě zahraničních (cizoměnových) transakcí jde o identifikátor transakce
type CisloDokladu struct {
}

func NewCisloDokladu() generator.Line {
	return &CisloDokladu{}
}

func (t *CisloDokladu) Generate(header model.Header, transfer model.Transaction) string {
	return transfer.CisloDokladu
}

func (t *CisloDokladu) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 13, "0", generator.SpacerLeft
}
