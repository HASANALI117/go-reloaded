package reboot

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// NumberConversion processes the input text and converts hexadecimal (hex) and binary (bin) numbers
// enclosed in parentheses to their decimal equivalents.
//
// The function searches for the following modifiers in the input text:
//   - (hex): Converts the hexadecimal number to decimal.
//   - (bin): Converts the binary number to decimal.
//
// The function is case-insensitive when matching modifiers.

func NumberConversion(text string) string {

	// Find all modifiers in the input text.
	modifiers := regexp.MustCompile(`(?i)\((\s+)?(hex|bin)(\s+)?\)`).FindAllString(text, -1)

	for _, modifier := range modifiers {

		// Extract the modifier type from the modifier string.
		modifierType := regexp.MustCompile(`(?i)(hex|bin)`).FindString(modifier)

		// Replace hexadecimal and binary numbers based on modifiers in the original text.
		text = regexp.MustCompile(`(\S+\s+)\((\s+)?`+modifierType+`(\s+)?\)`).ReplaceAllStringFunc(text, func(matchString string) string {

			// Extract the hexadecimal or binary number to be converted.
			modifiedString := regexp.MustCompile(`(^\S+[^\s+(]+)`).FindString(matchString)

			if !regexp.MustCompile(`^[0-9A-Fa-f]+$`).MatchString(modifiedString) {
				fmt.Printf("Add a number instead of: %s\n", matchString)
				return matchString
			}

			// Convert the number to decimal based on the modifier type.
			if strings.ToLower(modifierType) == "hex" {
				modifiedString = Hex2Decimal(modifiedString)
			} else if strings.ToLower(modifierType) == "bin" {
				modifiedString = Bin2Decimal(modifiedString)
			}

			return modifiedString
		})
	}
	return text
}

// Hex2Decimal converts a hexadecimal number to its decimal equivalent.
func Hex2Decimal(hexNum string) string {
	decNum, err := strconv.ParseInt(hexNum, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprint(decNum)
}

// Bin2Decimal converts a binary number to its decimal equivalent.
func Bin2Decimal(binNum string) string {
	decNum, err := strconv.ParseInt(binNum, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprint(decNum)
}
