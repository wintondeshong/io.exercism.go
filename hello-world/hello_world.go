package hello

import "fmt"

const testVersion = 2

// HelloWorld constructs a greeting to the caller. Attempts to return
// a personalized messaged message when 'name' is supplied. Otherwise,
// returns generalized greeting.
func HelloWorld(name string) string {
	if len(name) == 0 {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
