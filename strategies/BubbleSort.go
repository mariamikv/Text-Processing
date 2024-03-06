package strategies

type BubbleSort struct{}

func (b *BubbleSort) ProcessText(text string) (string, error) {
	runes := []rune(text)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes), nil
}
