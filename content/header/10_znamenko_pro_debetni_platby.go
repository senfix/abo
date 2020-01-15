package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Znaménko pro debetní obraty – vždy text 0
type ZnamenkoProDebetniPlatby struct {
}

func NewZnamenkoProDebetniPlatby() generator.Header {
	return &ZnamenkoProDebetniPlatby{}
}

func (t *ZnamenkoProDebetniPlatby) Generate(export model.Export) string {
	return "0"
}

func (t *ZnamenkoProDebetniPlatby) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 1, "0", generator.SpacerLeft
}
