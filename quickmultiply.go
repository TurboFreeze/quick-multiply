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

	data_record := ""

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
			// End timer
			end_time := time.Now()
			elapsed := end_time.Sub(start_time)

			fmt.Printf("%sSKIPPING%s due to invalid input, must be integer. It took %.2f seconds.\n", ANSI_YELLOW, ANSI_END, elapsed.Seconds())
			data_record += "0," + strconv.Itoa(num_first) + "," + strconv.Itoa(num_second) + "," + strconv.Itoa(prod) + ",," + strconv.FormatFloat(elapsed.Seconds(), 'f', 3, 64) // 0 = SKIP
		} else {
			// End timer
			end_time := time.Now()
			elapsed := end_time.Sub(start_time)

			// Check answer
			if prod == ans_int {
				fmt.Printf("%sCORRECT%s. It took %.2f seconds.\n", ANSI_GREEN, ANSI_END, elapsed.Seconds())
				data_record += "2," + strconv.Itoa(num_first) + "," + strconv.Itoa(num_second) + "," + strconv.Itoa(prod) + "," + strconv.Itoa(ans_int) + "," + strconv.FormatFloat(elapsed.Seconds(), 'f', 3, 64) // 2 = CORRECT
			} else {
				fmt.Printf("%sINCORRECT%s. It took %.2f seconds. The actual answer is %d.\n", ANSI_RED, ANSI_END, elapsed.Seconds(), prod)
				data_record += "1," + strconv.Itoa(num_first) + "," + strconv.Itoa(num_second) + "," + strconv.Itoa(prod) + "," + strconv.Itoa(ans_int) + "," + strconv.FormatFloat(elapsed.Seconds(), 'f', 3, 64) // -1 = INCORRECT
			}
		}
		fmt.Println(data_record)
		// Write the data
		// Schema is CSV: Status, first, second, product (real answer), input (user answer), time

		fmt.Println("*****************************")
	}

}
