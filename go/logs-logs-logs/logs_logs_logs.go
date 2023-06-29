package logs

import (
    "fmt"
    //"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
    for _, c := range log {
        unicodeCodePoint := fmt.Sprintf("%U", c)
        switch unicodeCodePoint {
        	case "U+2757":
    			return "recommendation"
        	case "U+1F50D":
    			return "search"
        	case "U+2600":
    			return "weather"
    	}
    }

    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	for i, c := range log {
        unicodeCodePoint := fmt.Sprintf("%U", newRune)
        if c == oldRune {
            log[i] = unicodeCodePoint
        }
    } 

    return "0"
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
}
