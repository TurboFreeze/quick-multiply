package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Continuous loop
	for true {
		// Numbers generated
		num_first := rand.Intn(99) + 1
		num_second := rand.Intn(99) + 1
		prod := num_first * num_second

		// Reading input
		fmt.Printf("%d x %d = ", num_first, num_second)

		// Begin timer
		start_time := time.Now()

		reader := bufio.NewReader(os.Stdin)
		ans, _ := reader.ReadString('\n')
		ans_int, err := strconv.Atoi(strings.TrimSpace(ans))
		if err != nil {
			fmt.Println("Invalid input, must be integer")
		} else {
			// End timer
			end_time := time.Now()
			elapsed := end_time.Sub(start_time)

			// Check answer
			if prod == ans_int {
				fmt.Printf("Correct. It took %.2f seconds.\n", elapsed.Seconds())
			} else {
				fmt.Printf("Incorrect. The actual answer is %d.\n", prod)
			}
		}
	}

}
