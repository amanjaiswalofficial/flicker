package session

import "fmt"

// Flicker initializes the flickerdfReader and prints a startup message.
// It returns an instance of the flickerdfReader.
func Flicker() flickerdfReader {
	fmt.Println("Starting up flicker...")
	f := flickerdfReader{}
	return f
}
