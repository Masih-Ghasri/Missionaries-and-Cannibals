package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("\n\tGame Start\nNow the task is to move all of them to the right side of the river")
	fmt.Println("Rules:\n1. The boat can carry at most two people\n2. If cannibals outnumber missionaries, the cannibals will eat the missionaries\n3. The boat cannot cross the river without people")

	lM, lC := 3, 3 // Left side Missionaries and Cannibals
	rM, rC := 0, 0 // Right side Missionaries and Cannibals
	attempts := 0

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Left to right travel
		for {
			fmt.Println("Left side -> Right side river travel")
			fmt.Print("Enter number of Missionaries to travel: ")
			scanner.Scan()
			uM, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Print("Enter number of Cannibals to travel: ")
			scanner.Scan()
			uC, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			if uM == 0 && uC == 0 {
				fmt.Println("Empty travel is not possible. Re-enter:")
			} else if (uM+uC) <= 2 && (lM-uM) >= 0 && (lC-uC) >= 0 {
				lM -= uM
				lC -= uC
				rM += uM
				rC += uC
				break
			} else {
				fmt.Println("Invalid input. Re-enter:")
			}
		}
		displayState(lM, lC, rM, rC, "->")
		attempts++

		if checkGameOver(lM, lC, rM, rC) {
			break
		}

		// Right to left travel
		for {
			fmt.Println("Right side -> Left side river travel")
			fmt.Print("Enter number of Missionaries to travel: ")
			scanner.Scan()
			userM, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Print("Enter number of Cannibals to travel: ")
			scanner.Scan()
			userC, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			if userM == 0 && userC == 0 {
				fmt.Println("Empty travel is not possible. Re-enter:")
			} else if (userM+userC) <= 2 && (rM-userM) >= 0 && (rC-userC) >= 0 {
				lM += userM
				lC += userC
				rM -= userM
				rC -= userC
				break
			} else {
				fmt.Println("Invalid input. Re-enter:")
			}
		}
		displayState(lM, lC, rM, rC, "<-")
		attempts++

		if checkGameOver(lM, lC, rM, rC) {
			break
		}
	}

	fmt.Printf("Total attempts: %d\n", attempts)
}

func displayState(lM, lC, rM, rC int, direction string) {
	fmt.Printf("\n")
	for i := 0; i < lM; i++ {
		fmt.Print("M ")
	}
	for i := 0; i < lC; i++ {
		fmt.Print("C ")
	}
	fmt.Printf("| %s | ", direction)
	for i := 0; i < rM; i++ {
		fmt.Print("M ")
	}
	for i := 0; i < rC; i++ {
		fmt.Print("C ")
	}
	fmt.Printf("\n")
}

func checkGameOver(lM, lC, rM, rC int) bool {
	if (lC > lM && lM > 0) || (rC > rM && rM > 0) {
		fmt.Println("Cannibals eat missionaries. You lost the game!")
		return true
	}
	if lM+lC == 0 {
		fmt.Println("You won the game. Congratulations!")
		return true
	}
	return false
}
