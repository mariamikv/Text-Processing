package strategies

type SelectionSort struct{}

func (s *SelectionSort) ProcessText(text string) (string, error) {
	runes := []rune(text)
	for i := 0; i < len(runes)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[minIndex] {
				minIndex = j
			}
		}
		runes[i], runes[minIndex] = runes[minIndex], runes[i]
	}
	return string(runes), nil
}
