package reboot

import "regexp"

func AllConversions(inputText string) string {
	// Apply a series of text transformations in a specific order.
	inputText = A2An(inputText)
	temp := ""
	for temp != inputText {
		temp = inputText
		inputText = Punct(inputText)
	}
	inputText = Punct2(inputText)
	inputText = NumberConversion(inputText)
	inputText = WordConversion(inputText)

	return inputText
}

// A2An replaces instances of "a" followed by a vowel with "an."
func A2An(str string) string {
	return regexp.MustCompile(`(?i)a\s+[aeiou]\w+`).ReplaceAllStringFunc(str, func(s string) string {
		// Replace "a" with "an" before words starting with a vowel.
		return regexp.MustCompile(`^[^\s+]`).ReplaceAllLiteralString(s, "an")
	})
}

// Punct removes extra spaces around punctuation marks.
func Punct(str string) string {
	return regexp.MustCompile(`(\s+)[.,!?:;]+(\s+)?`).ReplaceAllStringFunc(str, func(s string) string {
		// Remove extra spaces around punctuation marks and preserve the punctuation marks.
		sign := regexp.MustCompile(`[.,!?:;]+`).FindString(s)
		return sign + " "
	})
}

// Punct2 adjusts single quotes within text.
func Punct2(str string) string {
	return regexp.MustCompile(`'.*?'`).ReplaceAllStringFunc(str, func(s string) string {
		// Adjust single quotes within text.
		return regexp.MustCompile(`^('\s+)|(\s+')$`).ReplaceAllLiteralString(s, "'")
	})
}
