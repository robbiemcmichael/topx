package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s number file\n", os.Args[0])
		os.Exit(1)
	}

	err := run(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(number string, file string) error {
	top := &MinHeap{}
	heap.Init(top)

	x, err := strconv.Atoi(number)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return err
		}

		// Maintain a min-heap of the x largest numbers
		if len(*top) < x {
			heap.Push(top, n)
		} else if x > 0 && n > (*top)[0] {
			heap.Pop(top)
			heap.Push(top, n)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Reverse the ordering for output
	result := make([]float64, top.Len())
	for top.Len() > 0 {
		result[top.Len()-1] = heap.Pop(top).(float64)
	}

	for _, n := range result {
		fmt.Println(n)
	}

	return nil
}
