package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputs := getInputs()
	result := find2020(inputs)

	fmt.Println(result)

}

func find2020(inputs []int64) int64 {
	for _, i := range inputs {
		for _, h := range inputs {
			if (i + h) == 2020 {
				return i * h
			}
		}
	}

	panic("not found")
}

func getInputs() []int64 {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(dir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	inputs := []int64{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		a, _ := strconv.ParseInt(s.Text(), 0, 64)
		inputs = append(inputs, a)
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return inputs
}
