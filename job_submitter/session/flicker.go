package session

import "fmt"

func Flicker() flickerdfReader {
	fmt.Println("Starting up flicker...")
	f := flickerdfReader{}
	return f
}