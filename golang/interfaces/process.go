// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // DataProcessor interface defines a method for processing data
// type DataProcessor interface {
// 	Process(data string) string
// }

// // UppercaseProcessor struct converts input data to uppercase
// type UppercaseProcessor struct{}

// // Process method for UppercaseProcessor
// func (u UppercaseProcessor) Process(data string) string {
// 	return strings.ToUpper(data)
// }

// // ReverseProcessor struct reverses the input data
// type ReverseProcessor struct{}

// // Process method for ReverseProcessor
// func (r ReverseProcessor) Process(data string) string {
// 	runes := []rune(data)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// // Function that accepts a DataProcessor interface and returns an UppercaseProcessor
// func GetDataProcessor() DataProcessor {
// 	return ReverseProcessor{}
// }

// // Function that uses any DataProcessor to process input data
// func ProcessData(processor DataProcessor, data string) string {
// 	return processor.Process(data)
// }

// func main() {
// 	processor := GetDataProcessor()
// 	result := ProcessData(processor, "hello world")
// 	fmt.Println(result) // Output: HELLO WORLD
// }
