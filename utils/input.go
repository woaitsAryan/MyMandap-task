package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func GetInputString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatalln(error)
	}
	return strings.TrimSpace(input)
}

func GetInputInt(prompt string) int {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	integer, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Fatalln(err)
	}
    return integer
}