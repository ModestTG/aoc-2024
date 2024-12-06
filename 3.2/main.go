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
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	total := FindMuls(string(bytes))

	fmt.Printf("Total Sum: %v\n", total)
	FindDont(string(bytes))
	deduct := FindDont(string(bytes))
	fmt.Printf("New Sum: %v\n", total-deduct)
}

func FindMuls(s string) int {
	sum := 0
	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(s, -1)
	// fmt.Println(matches)
	for _, val := range matches {
		pd1 := `\d+`
		re1 := regexp.MustCompile(pd1)
		nums := re1.FindAllString(val, 2)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		sum += num1 * num2
	}
	return sum
}

func FindDont(s string) int {
	sum := 0
	pattern := `(?s)don't\(\)(.*?)do\(\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(s, -1)
	for _, val := range matches {
		// fmt.Printf("match: [[[ %v ]]]\n\n\n", val)
		sum += FindMuls(val)
	}
	return sum
}
