package main

import (
	"fmt"
	"strings"
)

type DataProcessor interface {
	Process(data string) string
}

type ReverseProcessor struct{}

func (r ReverseProcessor) Process(data string) string {
	runes := []rune(data)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func GetDataProcessor() DataProcessor {
	return ReverseProcessor{}
}

func ProcessData(processor DataProcessor, data string) string {
	return processor.Process(data)
}

func main() {
	processor := GetDataProcessor()
	result := ProcessData(processor, "hello world")
	fmt.Println(result)
}
