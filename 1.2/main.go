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
		count := 0
		for j := 0; j < len(sl2); j++ {
			if sl2[j] == sl1[i] {
				// fmt.Printf("%d matches %d", sl1[i], sl2[j])
				count++
			}
		}
		if count != 0 {
			fmt.Printf("%d | matched %d times\n", sl1[i], count)
		}
		sum += sl1[i] * count
		// fmt.Printf("1: %d | total similarity: %d\n", sl1[i], sum)
	}

	fmt.Printf("total diff: %d", sum)
}
