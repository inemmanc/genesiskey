package controller

import (
	"bufio"
	"fmt"
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
		// TEMP
		fmt.Println("DEFAULT STD")
	case "url":
		// TEMP
		fmt.Println("DEFAULT URL")
	default:
		// TEMP
		fmt.Println("DEFAULT STD init")
	}
}
