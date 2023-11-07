package reboot

import (
	"fmt"
	"regexp"
	"strings"
)

// WordConversion performs word transformations in the input text based on specified modifiers.
// It scans the input text for the following modifiers enclosed in parentheses and applies the corresponding transformations:
//
//   - (up): Converts the word before it to uppercase.
//   - (low): Converts the word before it to lowercase.
//   - (cap): Converts the word before it to capitalized form.
//
// The function is case-insensitive when matching modifiers.

func WordConversion(text string) string {

	// Find all modifiers in the input text.
	modifiers := regexp.MustCompile(`(?i)\((\s+)?(up|low|cap)(\s+)?(,(\s+)?(\d+)(\s+)?)?\)`).FindAllString(text, -1)

	for _, modifier := range modifiers {

		// Extract the modifier type and count from the modifier string.
		modifierType := regexp.MustCompile(`(?i)(\w+[a-z])`).FindString(modifier)
		modifierCount := regexp.MustCompile(`([0-9]+)`).FindString(modifier)

		// If no count is specified, default to 1.
		if modifierCount == "" {
			modifierCount = "1"
		}

		// If count is 0, skip the replacement.
		if modifierCount == "0" {
			fmt.Println("Count cannot be 0")
			continue
		}

		// Replace words based on modifiers in the original text.
		text = regexp.MustCompile(`(\S+\s+){`+modifierCount+`}\((\s+)?`+modifierType+`(\s+)?(,(\s+)?`+modifierCount+`(\s+)?)?\)`).ReplaceAllStringFunc(text, func(matchString string) string {

			// Extract the words to be modified.
			modifiedString := regexp.MustCompile(`(^[^(]+)`).FindString(matchString)

			// Remove trailing whitespace.
			modifiedString = regexp.MustCompile(`\s+$`).ReplaceAllLiteralString(modifiedString, "")

			// Apply the specified transformation to the word.
			if strings.ToLower(modifierType) == "up" {
				modifiedString = strings.ToUpper(modifiedString)
			} else if strings.ToLower(modifierType) == "low" {
				modifiedString = strings.ToLower(modifiedString)
			} else if strings.ToLower(modifierType) == "cap" {
				modifiedString = strings.Title(modifiedString)
			}

			return modifiedString
		})
	}
	return text
}
