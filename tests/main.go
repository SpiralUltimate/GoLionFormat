package main

// Import LionFormat and fmt
import (
	// Import LionFormat/format as LionFormat
	LionFormat "github.com/SpiralUltimate/GoLionFormat/format"

	// Import fmt for printing/editing strings
	"fmt"
)

func main() {
	// Message to format
	const msg = "val is {d}"

	// Format a new message using LionFormat

	formattedMsg, err := LionFormat.Format(msg, []any{0})

	// Check for errors
	if err != nil {
		// Handle error
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println(formattedMsg)
}
