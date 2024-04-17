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
	case "raw":
		fmt.Println("\nYOUR RAW KEY:")
		fmt.Println(gen.Generate("raw"))
	default:
		fmt.Println("\nINVALID OPERATION")
		SelectMethod()
	}
}

func Recall() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(text.RecallText())
	scanner.Scan()
	selected := strings.ToLower(scanner.Text())

	if (selected == "yes") || (selected == "y") {
		SelectMethod()
	}
}
