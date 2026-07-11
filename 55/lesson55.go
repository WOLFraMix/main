package main

import (
	"fmt"
	"strings"
)

// Calls UserProfileToString with sample data and prints the result or any error.
func main() {
	result, err := UserProfileToString("John", 25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// UserProfileToString formats a user profile into a string.
// It validates the name and age parameters, returning an error if validation fails.
func UserProfileToString(name string, age int) (string, error) {
	// Validate that the name is not empty
	if name == "" {
		return "", fmt.Errorf("empty name")
		// Validate that the age is not negative
	} else if age < 0 {
		return "", fmt.Errorf("negative age")
		// Validate that the name does not consist solely of whitespace
	} else if strings.TrimSpace(name) == "" {
		return "", fmt.Errorf("name cannot contain only spaces")
		// Return the formatted profile string if all checks pass
	} else {
		return fmt.Sprintf("Имя человека: %s, возраст: %d.", name, age), nil
	}
}
