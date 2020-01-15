package transaction

import (
	"fmt"

	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Číselný kód měny
//Používá se následující číselné označení (dle měny abecedně): 00036 pro AUD, 00124
//pro CAD, 00156 pro CNY, 00203 pro CZK, 00208 pro DKK, 00978 pro EUR, 00826
//pro GBP, 00191 pro HRK, 00348 pro HUF, 00756 pro CHF, 00392 pro JPY, 00578 pro
//NOK, 00985 pro PLN, 00946 pro RON, 00643 pro RUB, 00752 pro SEK, 00949 pro TRY,
//00840 pro USD
type KodMeny struct {
}

func NewKodMeny() generator.Line {
	return &KodMeny{}
}

func (t *KodMeny) Generate(header model.Header, transfer model.Transaction) string {
	return fmt.Sprintf("%v", transfer.KodMeny)
}

func (t *KodMeny) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 5, "0", generator.SpacerLeft
}
