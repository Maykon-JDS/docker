package main
import (
	"fmt"
)

func main() {
	fmt.Println("Qual o seu nome:?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Oi, %s! Eu sou a linguagem GO! ", name)
}