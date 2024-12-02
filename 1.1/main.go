package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sl1, sl2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val1, err := strconv.Atoi(strings.Split(scanner.Text(), "   ")[0])
		if err != nil {
			log.Fatal(err)
		}
		val2, err := strconv.Atoi(strings.Split(scanner.Text(), "   ")[1])
		if err != nil {
			log.Fatal(err)
		}
		sl1 = append(sl1, val1)
		sl2 = append(sl2, val2)
	}
	sort.Ints(sl1)
	sort.Ints(sl2)

	sum := 0
	for i := 0; i < len(sl1); i++ {
		diff := sl2[i] - sl1[i]
		if sl1[i] > sl2[i] {
			diff = -diff
		}
		sum += diff
		fmt.Printf("1: %d | 2: %d | total sum: %d\n", sl1[i], sl2[i], sum)

	}
	fmt.Printf("total diff: %d", sum)
}
