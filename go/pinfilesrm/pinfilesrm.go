// pinfilesrm
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		command := exec.Command("ipfs", "pin", "rm", line)
		out, cmderr := command.Output()
		if out != nil {
			fmt.Println(string(out))
		}
		if cmderr != nil {
			log.Println(cmderr)
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
