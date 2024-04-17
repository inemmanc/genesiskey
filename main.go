package main

import (
	"fmt"
	"genesiskey/src/controller"
	"genesiskey/src/text"
)

func main() {
	fmt.Println(text.DefaultWelcomeMessage())
	controller.SelectMethod()
	controller.Recall()
	fmt.Println(text.FinalScreenText())
	fmt.Scanln()
}
