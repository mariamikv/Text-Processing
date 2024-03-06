package context

import (
	"fmt"
	"main/textProcessor"
)

type TextProcessingContext struct {
	processor textProcessor.TextProcessor
}

func (ctx *TextProcessingContext) SetStrategy(processor textProcessor.TextProcessor) {
	ctx.processor = processor
}

func (ctx *TextProcessingContext) Process(text string) (string, error) {
	if ctx.processor == nil {
		return "", fmt.Errorf("no strategy set")
	}
	return ctx.processor.ProcessText(text)
}
