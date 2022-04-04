package main

import (
	. "fmt"
	"os"
	"strings"
)

var userinput string
var arrayuserinput []string

func main() {
	Scan(&userinput)
	stringtoarray(userinput)
	checkForPalin()
}
func stringtoarray(userinput string) {
	arrayuserinput = strings.Split(userinput, "")
}

//ik heb de namen logischer beschreven, de start is de index van de linkerzijde van een woord en 'einde'is de rechterzijde van een woord
func checkForPalin() {
	var start = 0
	var einde int
	start = len(arrayuserinput) - 1
	for einde > start {
		if arrayuserinput[start] != arrayuserinput[einde] {
			println("not a palindroom")
			os.Exit(0)
		} else {
			start++
			einde--
		}
	}
	println("is a palindroom")
}
