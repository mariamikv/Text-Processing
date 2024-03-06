package textProcessor

type TextProcessor interface {
	ProcessText(text string) (string, error)
}
