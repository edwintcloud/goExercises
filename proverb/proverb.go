// Package proverb is a package
package proverb

import "fmt"

// Proverb is a function
func Proverb(rhyme []string) []string {

	// if len rhyme is 0, return empty string slice
	if len(rhyme) == 0 {
		return []string{}
	}

	// build result
	result := []string{}
	for i := 0; i < len(rhyme)-1; i++ {
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	// return result
	return result
}
