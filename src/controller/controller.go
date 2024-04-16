package controller

import (
	"bufio"
	"fmt"
	"genesiskey/src/text"
	"os"
)

func SelectMethod() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(text.DefaultMethodText())
	scanner.Scan()
}
