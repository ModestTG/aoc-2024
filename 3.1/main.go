package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	sum := 0
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	pattern := `mul\(\d+,\d+\)`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)
	// fmt.Println(matches)
	for _, val := range matches {
		pd1 := `\d+`
		re1 := regexp.MustCompile(pd1)
		nums := re1.FindAllString(val, 2)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		sum += num1 * num2
	}
	fmt.Printf("Total Sum: %v", sum)
}
