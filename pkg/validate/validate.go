package validate

import (
	"fmt"
)

type (
	// Errors are returned with a slice of all invalid fields.
	Errors struct {
		Errors []Error
	}

	// Error knows for a given field the error.
	Error struct {
		Field string
		Error error
	}
)

func (e Errors) Error() string {
	for _, err := range e.Errors {
		fmt.Println(err)
	}
	return fmt.Sprintf("there are %d validation errors", len(e.Errors))
}
