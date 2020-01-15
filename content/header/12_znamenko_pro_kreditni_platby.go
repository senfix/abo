package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Znaménko pro kreditní obraty – vždy text 0
type ZnamenkoProKreditniPlatby struct {
}

func NewZnamenkoProKreditniPlatby() generator.Header {
	return &ZnamenkoProKreditniPlatby{}
}

func (t *ZnamenkoProKreditniPlatby) Generate(export model.Export) string {
	return "0"
}

func (t *ZnamenkoProKreditniPlatby) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 1, "0", generator.SpacerLeft
}
