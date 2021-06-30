package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Run(expenseReport string) {
	expenses := getExpenses(expenseReport)

out:
	// part I
	for i, v := range expenses {
		for j, v2 := range expenses {
			if i == j {
				continue
			}
			if v+v2 == 2020 {
				fmt.Println(v * v2)
				break out
			}
		}
	}

	// part II

	for i, v := range expenses {
		for j, v2 := range expenses {
			if i == j {
				continue
			}
			for k, v3 := range expenses {
				if k == i || k == j {
					continue
				}
				if v+v2+v3 == 2020 {
					fmt.Println(v * v2 * v3)
					return
				}

			}
		}
	}
}

func getExpenses(iFile string) []int {
	var expenses []int
	file, err := os.Open(iFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		expenses = append(expenses, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return expenses
}
