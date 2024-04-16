package controller

import (
	"bufio"
	"fmt"
	"genesiskey/src/gen"
	"genesiskey/src/text"
	"os"
	"strings"
)

func SelectMethod() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(text.DefaultMethodText())
	scanner.Scan()
	selected := strings.ToLower(scanner.Text())

	switch selected {
	case "std":
		fmt.Println(gen.Generate("test"))
	case "url":
		// TEMP
		fmt.Println("DEFAULT URL")
	default:
		// TEMP
		fmt.Println("\nINVALID OPERATION")
		fmt.Println()
		SelectMethod()
	}
}
