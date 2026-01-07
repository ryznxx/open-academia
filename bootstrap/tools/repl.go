// code written github.com/ryznxx
package tools

import (
	"bootstrap-open-academia/target"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Repl() {
	availableDevice := []string{"Windows", "Linux", "BSD"}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n\n============ Open-Academia Installer ===========\n\n")

	for i := 0; i < len(availableDevice); i++ {
		fmt.Printf("%d. %s\n", i+1, availableDevice[i])
	}

loop:
	for {
		fmt.Print("\nDevice Target : ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		switch strings.ToLower(answer) {
		case "1", strings.ToLower("Windows"):
			fmt.Print("windows")
			break loop
		case "2", strings.ToLower("Linux"):
			target.Linux()
			break loop
		case "3", strings.ToLower("BSD"):
			fmt.Print("bsd")
			break loop
		default:
			fmt.Println("salah woi")
		}
	}

}
