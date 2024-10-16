package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	summer_array []int
	summer       int
	flag1        bool
)

func main() {
	file, err := os.Open("main2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	now := time.Now()
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		go divi(words)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(summer_array)
	fmt.Println("Время выполнения:", time.Since(now))
	fmt.Scan(flag1)
}
func divi(words []string) {
	for _, word := range words {
		price, err := strconv.Atoi(word)
		if err != nil {
			panic(err)
		}
		summer += price
	}
	summer_array = append(summer_array, summer/len(words))
	summer = 0
}
