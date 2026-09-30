package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	firstRun := true

	loopApp(scanner, firstRun)
}
