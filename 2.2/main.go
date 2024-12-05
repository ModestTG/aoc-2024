package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var count int = 0

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var sl1 []int
		chars := strings.Split(scanner.Text(), " ")
		for _, val := range chars {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			sl1 = append(sl1, num)
		}
		if !EvalLevel(sl1) {
			loopWithOneRemoved(sl1)
		}
	}
	fmt.Printf("Total safe: %d", count)
}

func EvalLevel(level []int) bool {
	fmt.Printf("%d ", level)
	for i := range level {
		// considered safe if on last value
		if i+1 == len(level) {
			fmt.Println("safe")
			count++
			return true
		}
		// unsafe if values are not all increasing or decreasing
		if i > 0 {
			tmp1 := level[i-1] - level[i]
			tmp2 := level[i] - level[i+1]
			if (tmp1 > 0 && tmp2 < 0) || (tmp1 < 0 && tmp2 > 0) {
				fmt.Println("unsafe, values not all inc or dec")
				return false
			}
		}
		// unsafe if adjacent values are neither increasing or decreasing
		if level[i] == level[i+1] {
			fmt.Println("unsafe, adjacent values not inc or dec")
			return false
		}
		tmp3 := level[i] - level[i+1]
		// unsafe if values differ too much
		if tmp3 > 3 || tmp3 < -3 {
			fmt.Println("unsafe, adjacent values vary by more than 3")
			return false
		}
	}
	return true
}

func loopWithOneRemoved(slice []int) {
	for i := 0; i < len(slice); i++ {
		// fmt.Printf("Removing value: %d\n", slice[i])

		// Create a new slice excluding the current element
		newSlice := append([]int{}, slice[:i]...)
		newSlice = append(newSlice, slice[i+1:]...)
		// fmt.Println("Remaining slice:", newSlice)

		if EvalLevel(newSlice) {
			break
		}
	}
}
