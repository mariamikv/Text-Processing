package strategies

type InsertionSort struct{}

func (s *InsertionSort) ProcessText(text string) (string, error) {
	runes := []rune(text)
	for i := 1; i < len(runes); i++ {
		key := runes[i]
		j := i - 1
		for j >= 0 && key < runes[j] {
			runes[j+1] = runes[j]
			j--
		}
		runes[j+1] = key
	}
	return string(runes), nil
}
