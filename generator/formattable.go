package generator

type Formattable interface {
	Restrictions() (length int, spacer string, position SpacerPosition)
}
