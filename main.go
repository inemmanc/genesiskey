package main

import (
	"fmt"
	"genesiskey/src/controller"
	"genesiskey/src/text"
)

func main() {
	fmt.Println(text.DefaultWelcomeMessage())
	controller.SelectMethod()
}
