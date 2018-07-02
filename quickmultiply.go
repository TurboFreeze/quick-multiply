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
	// ANSI sequences
	ANSI_END := "\033[0m"
	ANSI_RED := "\033[31m"
	ANSI_GREEN := "\033[32m"
	ANSI_YELLOW := "\033[33m"

	// Seed RNG
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("*****************************")
	fmt.Println("* WELCOME TO QUICK MULTIPLY *")
	fmt.Println("*****************************")

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
			fmt.Printf("%sSKIPPING%s due to invalid input, must be integer.\n", ANSI_YELLOW, ANSI_END)
		} else {
			// End timer
			end_time := time.Now()
			elapsed := end_time.Sub(start_time)

			// Check answer
			if prod == ans_int {
				fmt.Printf("%sCORRECT%s. It took %.2f seconds.\n", ANSI_GREEN, ANSI_END, elapsed.Seconds())
			} else {
				fmt.Printf("%sINCORRECT%s. The actual answer is %d.\n", ANSI_RED, ANSI_END, prod)
			}
		}

		fmt.Println("*****************************")
	}

}
