package golangexamples

import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string {
	var RetString string
	for i := 0; i < len(sliceToConcat); i++ {
		RetString += string(sliceToConcat[i]) + string('-')
	}
	RetString += string('\n')
	return RetString
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i := 0; i < len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] += byte(ceaserCount)
	}
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
