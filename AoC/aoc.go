package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	text := string(fileContent)

	// process string
	lines := strings.Split(text, "\n")

	var num int
	var biggest int
	biggest = 0

	for _, element := range lines {
		element = strings.TrimSpace(element)
		if (element != "") {
			number, err := strconv.Atoi(element)

			if (err != nil) {
				fmt.Println("ERROR: ", err)
			}

			num += number
		} else {
			if (num > biggest) {
				biggest = num
			}
			// fmt.Println("SUM: ", num)
			num = 0
		}

	}

	fmt.Println("BIGGEST: ", biggest)
	// fmt.Println(text)
}