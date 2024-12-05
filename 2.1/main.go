package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
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
		fmt.Printf("%d ", sl1)
		for i := range sl1 {
			// considered safe if on last value
			if i+1 == len(sl1) {
				count++
				fmt.Printf("safe - count: %d\n", count)
				break
			}
			// unsafe if values are not all increasing or decreasing
			if i > 0 {
				tmp1 := sl1[i-1] - sl1[i]
				tmp2 := sl1[i] - sl1[i+1]
				if (tmp1 > 0 && tmp2 < 0) || (tmp1 < 0 && tmp2 > 0) {
					fmt.Println("unsafe, values not all inc or dec")
					break
				}
			}
			// unsafe if adjacent values are neither increasing or decreasing
			if sl1[i] == sl1[i+1] {
				fmt.Println("unsafe, adjacent values not inc or dec")
				break
			}
			tmp3 := sl1[i] - sl1[i+1]
			// unsafe if values differ too much
			if tmp3 > 3 || tmp3 < -3 {
				fmt.Println("unsafe, adjacent values vary by more than 3")
				break
			}
		}
	}
	fmt.Printf("Total safe: %d", count)
}
