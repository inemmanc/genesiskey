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
		fmt.Println("\nYOUR STANDARD KEY:")
		fmt.Println(gen.Generate("std"))
	case "url":
		fmt.Println("\nYOUR URL KEY:")
		fmt.Println(gen.Generate("url"))
	default:
		fmt.Println("\nINVALID OPERATION")
		SelectMethod()
	}
}
