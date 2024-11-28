// This file defines functions similar to fmt.Printf
package format

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// reflect.Type of different types
var stringType = reflect.TypeOf("")
var intType = reflect.TypeOf(int(0))
var float64Type = reflect.TypeOf(float64(0.0))

// Checks if checkType == type of value
func isType(value any, checkType reflect.Type) bool {
	return reflect.TypeOf(value).AssignableTo(checkType)
}

// Generates a format type error, given the type (string) and value (any)
func formatTypeError(typeStr string, value any) error {
	return fmt.Errorf("format symbol of type '%v' doesn't match value '%v'", typeStr, value)
}

// Formats a string using the given format specifier and arguments.
func Format(format string, args ...any) (string, error) {
	// An array of runes (chars) representing the format
	formatRunes := []rune(format)

	// Collect all format symbols (eg. {s}, {d}, {any})
	for i := range args {
		// Get format symbol
		var formatSymbol string = string(formatRunes[strings.Index(format, "{") : strings.Index(format, "}")+1])
		// The arg corresponding to the format symbol
		var formatArg any = args[i]

		// Type check formatSymbol to ensure correct format specifier
		switch formatSymbol {
		// String
		case "{s}":
			if !isType(formatArg, stringType) {
				// Error message to display
				var errMessage = formatTypeError(stringType.String(), formatArg)

				return format, errMessage
			}
		// Integers/Float
		case "{d}":
			if !(isType(formatArg, intType) && !isType(formatArg, float64Type)) {
				var errMessage = formatTypeError("int | float64", formatArg)
				return format, errMessage
			}

		// Cases where format symbol is not recognized
		default:
			return format, errors.New("Invalid format type for format symbol " + formatSymbol)
		}

		// Replace format specifier
		format = strings.Replace(format, formatSymbol, fmt.Sprintf("%v", formatArg), 1)

		// Reassign format runes to keep up with the format string
		formatRunes = []rune(format)
	}

	return format, nil
}
