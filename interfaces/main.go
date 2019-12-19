package main

import "fmt"

type bot interface {
	getGreeting() string //any type that matches this description or is using this function with return type string will automatically be upgraded to type bot
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//VERY custom logic for generating english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
