package main

import (
	"fmt"
	"main/context"
	"main/strategies"
	"main/textProcessor"
)

func main() {
	text := "Text Processing Project : date 2024.03.07"

	textProcessingContext := context.TextProcessingContext{}

	textProcessingStrategies := []struct {
		name     string
		strategy textProcessor.TextProcessor
	}{
		{"Bubble Sort", &strategies.BubbleSort{}},
		{"Insertion Sort", &strategies.InsertionSort{}},
		{"Selection Sort", &strategies.SelectionSort{}},
		{"Regex Filter (numbers)", &strategies.RegexFilter{Pattern: "\\d"}},
	}

	for _, strategy := range textProcessingStrategies {
		textProcessingContext.SetStrategy(strategy.strategy)

		processedText, err := textProcessingContext.Process(text)
		if err != nil {
			fmt.Println("Error processing text with", strategy.name, ":", err)
			continue
		}

		fmt.Println("Result using", strategy.name, ":")
		fmt.Println(processedText)
		fmt.Println("----")
	}
}
